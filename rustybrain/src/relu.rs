pub struct ActivationReLu {
    pub outputs : Vec<f32>
}


impl ActivationReLu {
    pub fn new() -> ActivationReLu {
        return ActivationReLu {outputs: Vec::new()};
    }

    pub fn forward(&mut self, inputs: Vec<f32>) {
        for i in inputs.iter() {
            if *i < 0.0 {
                self.outputs.push(0.0);
            } else {
                self.outputs.push(*i);
            }
        }
    }
}
