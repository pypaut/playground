use bevy::{prelude::*, sprite::MaterialMesh2dBundle};

pub struct Player {
    color: Color,
    size: Vec2,
    transform: Transform,
}

impl Player {
    pub fn new(color: Color, size: Vec2, transform: Transform) -> Player {
        return Player {
            color: color,
            size: size,
            transform: transform,
        };
    }

    pub fn spawn(self, commands: &mut Commands) {
        commands.spawn(SpriteBundle {
            sprite: Sprite {
                color: self.color,
                custom_size: Some(self.size),
                ..default()
            },
            transform: self.transform,
            ..default()
        });
    }
}
