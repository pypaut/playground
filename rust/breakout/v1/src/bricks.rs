use bevy::{prelude::*, sprite::MaterialMesh2dBundle};

pub struct Brick {
    color: Color,
    size: Vec2,
    transform: Transform,
}

pub struct Bricks {
    bricks: Vec<Brick>,
}

impl Brick {
    pub fn new(color: Color, size: Vec2, transform: Transform) -> Brick {
        return Brick {
            color: color,
            size: size,
            transform: transform,
        };
    }

    pub fn spawn(&self, commands: &mut Commands) {
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

impl Bricks {
    pub fn new(
        color: Color,
        size: Vec2,
        start_x: f32,
        start_y: f32,
        brick_gap: f32,
        nb_lines: i32,
        nb_columns: i32,
    ) -> Bricks {
        let mut bricks = vec![];

        for l in 0..nb_lines {
            for c in 0..nb_columns {
                let transform_x = start_x + c as f32 * (size.x + brick_gap);
                let transform_y = start_y - l as f32 * (size.y + brick_gap);

                bricks.push(Brick::new(
                    color,
                    size,
                    Transform::from_xyz(transform_x, transform_y, 0.),
                ));
            }
        }

        return Bricks { bricks: bricks };
    }

    pub fn spawn(self, commands: &mut Commands) {
        for b in self.bricks.iter() {
            b.spawn(commands);
        }
    }
}
