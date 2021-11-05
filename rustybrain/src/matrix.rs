use std::ops;
use std::vec::Vec;


#[derive(Debug)]
pub struct Matrix {
    pub n : usize,
    pub p : usize,
    pub val : Vec<Vec<f32>>
}


impl Matrix {
    pub fn new(n: usize, p: usize) -> Matrix {
        let mut val = vec![];
        for i in 0..n {
            val.push(vec![]);
            for j in 0..p {
                val[i].push(0.0);
            }
        }
        return Matrix {n: n, p: p, val: val};
    }

    pub fn mult(&self, rhs : &Matrix) -> Matrix {
        if self.p != rhs.n {
            panic!("Mind the dimension!");
        }
        let mut val = Vec::new();
        for i in 0..self.n {
            let mut line = Vec::new();
            for j in 0..rhs.p {
                let mut e = 0.0;
                for k in 0..self.p {
                    e = e + self.val[i][k] * rhs.val[k][j]
                }
                line.push(e);
            }
            val.push(line)
        }
        return Matrix {n: self.n, p: rhs.p, val: val};
    }

    pub fn transposed(&self) -> Matrix {
        let mut val = Vec::new();
        for j in 0..self.p {
            let mut line = Vec::new();
            for i in 0..self.n {
                line.push(self.val[i][j]);
            }
            val.push(line);
        }
        return Matrix {n: self.p, p: self.n, val: val};
    }
}

impl ops::Add<&Matrix> for &Matrix {
    type Output = Matrix;

    fn add(self, rhs: &Matrix) -> Matrix {
        if self.n != rhs.n || self.p != rhs.p {
            panic!("Mind the dimension!");
        }
        let mut val = Vec::new();
        for i in 0..self.n {
            let mut line = Vec::new();
            for j in 0..self.p {
                line.push(self.val[i][j] + rhs.val[i][j]);
            }
            val.push(line);
        }
        return Matrix { n: self.n, p: self.p, val: val};
    }
}
