use crate::matrix::Matrix;
use rand::prelude::*;

mod matrix;


pub struct DataGenerator {
}

impl DataGenerator {
    pub fn generate(nb_points: usize, nb_classes: usize) -> Matrix, Matrix {
        X = Matrix::new(nb_points*nb_classes, 2);
        y = Matrix::new(nb_points*nb_classes, 1);
        let mut rng = rand::thread_rng();
        for i in 0..nb_classes {
            let ix = (nb_points*i...nb_points*(i+1)).collect::<Vec<f32>>();
            let r = (0.0..1.0).step_by(1.0/nb_points).collect::<Vec<f32>>();
            let t = (i*4..(i+1)*4).step_by(1.0/nb_points).collect::<Vec<f32>>();
            for j in 0..nb_points {
                t[j] = t[j] + rng.gen() * 0.2;
            }
            let mut new_row = Matrix::new(2,1);
            // new_row.val = vec![r *];
            // for
            // X[ix] = ;
            // y[ix] = i;
        }
    }
}

