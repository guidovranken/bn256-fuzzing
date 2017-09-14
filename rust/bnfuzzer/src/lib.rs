#![feature(libc)]
extern crate bn;
extern crate libc;
extern crate bigint;
#[macro_use]
extern crate lazy_static;
use std::io;
use std::io::Read;
use std::slice;
pub use bigint::U256;
use libc::{uint8_t, size_t};
use std::cmp::min;
use std::ops::{Deref, DerefMut};
use std::ptr::copy_nonoverlapping;

#[derive(Debug)]
pub struct Error(pub &'static str);

impl From<&'static str> for Error {
    fn from(val: &'static str) -> Self {
        Error(val)
    }
}

pub trait Impl: Send + Sync {
	/// execute this built-in on the given input, writing to the given output.
	fn execute(&self, input: &[u8], output: &mut BytesRef) -> Result<(), Error>;
}

pub enum BytesRef<'a> {
	/// This is a reference to a vector
	Flexible(&'a mut Bytes),
	/// This is a reference to a slice
	Fixed(&'a mut [u8])
}

impl<'a> BytesRef<'a> {
	/// Writes given `input` to this `BytesRef` starting at `offset`.
	/// Returns number of bytes written to the ref.
	/// NOTE can return number greater then `input.len()` in case flexible vector had to be extended.
	pub fn write(&mut self, offset: usize, input: &[u8]) -> usize {
		match *self {
			BytesRef::Flexible(ref mut data) => {
				let data_len = data.len();
				let wrote = input.len() + if data_len > offset { 0 } else { offset - data_len };

				data.resize(offset, 0);
				data.extend_from_slice(input);
				wrote
			},
			BytesRef::Fixed(ref mut data) if offset < data.len() => {
				let max = min(data.len() - offset, input.len());
				for i in 0..max {
					data[offset + i] = input[i];
				}
				max
			},
			_ => 0
		}
	}
}

impl<'a> Deref for BytesRef<'a> {
	type Target = [u8];

	fn deref(&self) -> &[u8] {
		match *self {
			BytesRef::Flexible(ref bytes) => bytes,
			BytesRef::Fixed(ref bytes) => bytes,
		}
	}
}

impl <'a> DerefMut for BytesRef<'a> {
	fn deref_mut(&mut self) -> &mut [u8] {
		match *self {
			BytesRef::Flexible(ref mut bytes) => bytes,
			BytesRef::Fixed(ref mut bytes) => bytes,
		}
	}
}

/// Vector of bytes.
pub type Bytes = Vec<u8>;

fn read_point(reader: &mut io::Chain<&[u8], io::Repeat>) -> Result<::bn::G1, Error> {
	use bn::{Fq, AffineG1, G1, Group};

	let mut buf = [0u8; 32];

	reader.read_exact(&mut buf[..]).expect("reading from zero-extended memory cannot fail; qed");
    let px = Fq::from_slice(&buf[0..32]).map_err(|_| Error::from("Invalid point x coordinate"))?;

	reader.read_exact(&mut buf[..]).expect("reading from zero-extended memory cannot fail; qed");
	let py = Fq::from_slice(&buf[0..32]).map_err(|_| Error::from("Invalid point y coordinate"))?;

	Ok(
		if px == Fq::zero() && py == Fq::zero() {
			G1::zero()
		} else {
			AffineG1::new(px, py).map_err(|_| Error::from("Invalid curve point"))?.into()
		}
	)
}

fn read_fr(reader: &mut io::Chain<&[u8], io::Repeat>) -> Result<::bn::Fr, Error> {
	let mut buf = [0u8; 32];

	reader.read_exact(&mut buf[..]).expect("reading from zero-extended memory cannot fail; qed");
	::bn::Fr::from_slice(&buf[0..32]).map_err(|_| Error::from("Invalid field element"))
}

impl Impl for Bn128AddImpl {
	// Can fail if any of the 2 points does not belong the bn128 curve
	fn execute(&self, input: &[u8], output: &mut BytesRef) -> Result<(), Error> {
		use bn::AffineG1;

		let mut padded_input = input.chain(io::repeat(0));
		let p1 = read_point(&mut padded_input)?;
		let p2 = read_point(&mut padded_input)?;

		let mut write_buf = [0u8; 64];
		if let Some(sum) = AffineG1::from_jacobian(p1 + p2) {
			// point not at infinity
			sum.x().to_big_endian(&mut write_buf[0..32]).expect("Cannot fail since 0..32 is 32-byte length");
			sum.y().to_big_endian(&mut write_buf[32..64]).expect("Cannot fail since 32..64 is 32-byte length");;
		}
		output.write(0, &write_buf);

		Ok(())
	}
}

impl Impl for Bn128MulImpl {
	// Can fail if first paramter (bn128 curve point) does not actually belong to the curve
	fn execute(&self, input: &[u8], output: &mut BytesRef) -> Result<(), Error> {
		use bn::AffineG1;

		let mut padded_input = input.chain(io::repeat(0));
		let p = read_point(&mut padded_input)?;
		let fr = read_fr(&mut padded_input)?;

		let mut write_buf = [0u8; 64];
		if let Some(sum) = AffineG1::from_jacobian(p * fr) {
			// point not at infinity
			sum.x().to_big_endian(&mut write_buf[0..32]).expect("Cannot fail since 0..32 is 32-byte length");
			sum.y().to_big_endian(&mut write_buf[32..64]).expect("Cannot fail since 32..64 is 32-byte length");;
		}
		output.write(0, &write_buf);
		Ok(())
	}
}

mod bn128_gen {
	use bn::{AffineG1, AffineG2, Fq, Fq2, G1, G2, Gt, pairing};

	lazy_static! {
		pub static ref P1: G1 = G1::from(AffineG1::new(
			Fq::from_str("1").expect("1 is a valid field element"),
			Fq::from_str("2").expect("2 is a valid field element"),
		).expect("Generator P1(1, 2) is a valid curve point"));
	}

	lazy_static! {
		pub static ref P2: G2 = G2::from(AffineG2::new(
			Fq2::new(
				Fq::from_str("10857046999023057135944570762232829481370756359578518086990519993285655852781")
					.expect("a valid field element"),
				Fq::from_str("11559732032986387107991004021392285783925812861821192530917403151452391805634")
					.expect("a valid field element"),
			),
			Fq2::new(
				Fq::from_str("8495653923123431417604973247489272438418190587263600148770280649306958101930")
					.expect("a valid field element"),
				Fq::from_str("4082367875863433681332203403145435568316851327593401208105741076214120093531")
					.expect("a valid field element"),
			),
		).expect("the generator P2(10857046999023057135944570762232829481370756359578518086990519993285655852781 + 11559732032986387107991004021392285783925812861821192530917403151452391805634i, 8495653923123431417604973247489272438418190587263600148770280649306958101930 + 4082367875863433681332203403145435568316851327593401208105741076214120093531i) is a valid curve point"));
	}

	lazy_static! {
		pub static ref P1_P2_PAIRING: Gt = pairing(P1.clone(), P2.clone());
	}
}

impl Impl for Bn128PairingImpl {
	/// Can fail if:
	///     - input length is not a multiple of 192
	///     - any of odd points does not belong to bn128 curve
	///     - any of even points does not belong to the twisted bn128 curve over the field F_p^2 = F_p[i] / (i^2 + 1)
	fn execute(&self, input: &[u8], output: &mut BytesRef) -> Result<(), Error> {
		use bn::{AffineG1, AffineG2, Fq, Fq2, pairing, G1, G2, Gt};

		let elements = input.len() / 192; // (a, b_a, b_b - each 64-byte affine coordinates)
		if input.len() % 192 != 0 {
			return Err("Invalid input length, must be multiple of 192 (3 * (32*2))".into())
		}
		let ret_val = if input.len() == 0 {
			U256::one()
		} else {
			let mut vals = Vec::new();
			for idx in 0..elements {
				let a_x = Fq::from_slice(&input[idx*192..idx*192+32])
					.map_err(|_| Error::from("Invalid a argument x coordinate"))?;

				let a_y = Fq::from_slice(&input[idx*192+32..idx*192+64])
					.map_err(|_| Error::from("Invalid a argument y coordinate"))?;

				let b_b_x = Fq::from_slice(&input[idx*192+64..idx*192+96])
					.map_err(|_| Error::from("Invalid b argument imaginary coeff x coordinate"))?;

				let b_b_y = Fq::from_slice(&input[idx*192+96..idx*192+128])
					.map_err(|_| Error::from("Invalid b argument imaginary coeff y coordinate"))?;

				let b_a_x = Fq::from_slice(&input[idx*192+128..idx*192+160])
					.map_err(|_| Error::from("Invalid b argument real coeff x coordinate"))?;

				let b_a_y = Fq::from_slice(&input[idx*192+160..idx*192+192])
					.map_err(|_| Error::from("Invalid b argument real coeff y coordinate"))?;

				vals.push((
					G1::from(
						AffineG1::new(a_x, a_y).map_err(|_| Error::from("Invalid a argument - not on curve"))?
					),
					G2::from(
						AffineG2::new(
							Fq2::new(b_a_x, b_a_y),
							Fq2::new(b_b_x, b_b_y),
						).map_err(|_| Error::from("Invalid b argument - not on curve"))?
					),
				));
			};

			let mul = vals.into_iter().fold(Gt::one(), |s, (a, b)| s * pairing(a, b));

			if mul == *bn128_gen::P1_P2_PAIRING {
				U256::one()
			} else {
				U256::zero()
			}
		};

		let mut buf = [0u8; 32];
		ret_val.to_big_endian(&mut buf);
		output.write(0, &buf);

		Ok(())
	}
}

#[derive(Debug)]
struct Bn128AddImpl;

#[derive(Debug)]
struct Bn128MulImpl;

#[derive(Debug)]
struct Bn128PairingImpl;

// Ethereum builtin creator.
fn ethereum_builtin(name: &str) -> Box<Impl> {
	match name {
		"bn128_add" => Box::new(Bn128AddImpl) as Box<Impl>,
		"bn128_mul" => Box::new(Bn128MulImpl) as Box<Impl>,
		"bn128_pairing" => Box::new(Bn128PairingImpl) as Box<Impl>,
		_ => panic!("invalid builtin name: {}", name),
	}
}

#[no_mangle]
pub extern fn rustbnadd(data: *const uint8_t, len: size_t, output_c: *mut uint8_t) {
    let numbers = unsafe {
        assert!(!data.is_null());

        slice::from_raw_parts(data, len as usize)
    };
    let f = ethereum_builtin("bn128_add");
    let mut output = vec![0u8; 64];
    //f.execute(numbers, &mut BytesRef::Fixed(&mut output[..])).expect("Builtin should not fail");
    f.execute(numbers, &mut BytesRef::Fixed(&mut output[..])).unwrap_or( () );
    unsafe {
        copy_nonoverlapping(output.as_ptr(), output_c, output.len());
    }
}

#[no_mangle]
pub extern fn rustbnmul(data: *const uint8_t, len: size_t, output_c: *mut uint8_t) {
    let numbers = unsafe {
        assert!(!data.is_null());

        slice::from_raw_parts(data, len as usize)
    };
    let f = ethereum_builtin("bn128_mul");
    let mut output = vec![0u8; 64];
    //f.execute(numbers, &mut BytesRef::Fixed(&mut output[..])).expect("Builtin should not fail");
    f.execute(numbers, &mut BytesRef::Fixed(&mut output[..])).unwrap_or( () );
    unsafe {
        copy_nonoverlapping(output.as_ptr(), output_c, output.len());
    }
}

#[no_mangle]
pub extern fn rustbnpairing(data: *const uint8_t, len: size_t, output_c: *mut uint8_t) {
    let numbers = unsafe {
        assert!(!data.is_null());

        slice::from_raw_parts(data, len as usize)
    };
    let f = ethereum_builtin("bn128_pairing");
    let mut output = vec![0u8; 32];
    //f.execute(numbers, &mut BytesRef::Fixed(&mut output[..])).expect("Builtin should not fail");
    f.execute(numbers, &mut BytesRef::Fixed(&mut output[..])).unwrap_or( () );
    unsafe {
        copy_nonoverlapping(output.as_ptr(), output_c, output.len());
    }
}
