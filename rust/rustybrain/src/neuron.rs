pub struct Neuron<'a> {
    pub size : usize,
    pub weights : &'a Vec<f32>,
    pub bias : f32
}

impl Neuron<'_> {
    pub fn output(&self, input : &Vec<f32>) -> f32 {
        if input.len() != self.size {
            panic!("You fool.");
        }

        let mut output : f32 = 0f32;
        for i in 0..self.size {
            output += self.weights[i] * input[i];
        }

        return output + self.bias;
    }
}
