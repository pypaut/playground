// use crate::layer::Layer;
// use crate::matrix::Matrix;
use crate::relu::ActivationReLu;

// mod layer;
// mod matrix;
mod relu;


fn main() {

    // Inputs
    // let X = Matrix {
    //     n: 3,
    //     p: 4,
    //     val: vec![
    //         vec![1.0, 2.0, 3.0, 2.5],
    //         vec![2.0, 5.0, -1.0, 2.0],
    //         vec![-1.5, 2.7, 3.3, -0.8]
    //     ]
    // };

    // ReLU test
    let mut r = ActivationReLu::new();
    let inputs : Vec<f32> = vec![0.0, 2.0, -1.0, 3.3, -2.7, 1.1, 2.2, -100.0];
    r.forward(inputs);
    println!("Outputs: {:?}", r.outputs);

   //  // Initialize layer
   //  let l = Layer::new(3, 4);

   //  // Compute outputs
   //  let outputs = l.forward(X);

   //  // Display
   //  println!("Outputs: {?}", outputs.val);

}
