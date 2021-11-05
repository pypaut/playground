use crate::matrix::Matrix;
use rand::prelude::*;


pub struct Layer {
    n_inputs : usize,
    n_neurons : usize,
    weights : Matrix,
    biases : Matrix,
}

impl Layer {
    pub fn new(n_inputs: usize, n_neurons: usize) -> Layer {
        let mut rng = rand::thread_rng();

        // Weights
        let mut w = Matrix::new(n_inputs, n_neurons);
        for i in 0..n_inputs {
            for j in 0..n_neurons {
                w.val[i][j] = rng.gen();
            }
        }

        // Biases
        let mut b = Matrix::new(n_inputs, n_neurons);

        return Layer {
            n_inputs: n_inputs,
            n_neurons: n_neurons,
            weights: w,
            biases: b
        };
    }

    pub fn get_n_inputs(&self) -> usize {
        return self.n_inputs;
    }

    pub fn get_n_neurons(&self) -> usize {
        return self.n_neurons;
    }

    pub fn get_weights(&self) -> &Matrix {
        return &self.weights;
    }

    pub fn get_biases(&self) -> &Matrix {
        return &self.biases;
    }

    pub fn forward(&self, inputs: Matrix) -> Matrix {
        return &inputs.mult(&self.get_weights().transposed()) + &self.biases;
    }
}

