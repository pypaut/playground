use bevy::prelude::Component;

// region:    --- Common Components

#[derive(Component)]
pub struct Direction {
    pub x: f32,
    pub y: f32,
}

impl Direction {
    pub fn norm(&self) -> f32 {
        (self.x.powf(2.0) + self.y.powf(2.)).sqrt()
    }

    pub fn normalize(&mut self) {
        let norm = self.norm();
        self.x /= norm;
        self.y /= norm;
    }
}

// endregion: --- Common Components

// region:    --- Player Components

#[derive(Component)]
pub struct Player;

// endregion: --- Player Components

// region:    --- Ball Components

#[derive(Component)]
pub struct Ball;

// endregion: --- Ball Components

// region:    --- Brick Components

#[derive(Component)]
pub struct Brick;

// endregion: --- Brick Components
