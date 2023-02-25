use bevy::{prelude::*, sprite::MaterialMesh2dBundle};

pub struct Ball {
    transform: Transform,
}

impl Ball {
    pub fn new(
        commands: &mut Commands,
        meshes: &mut ResMut<Assets<Mesh>>,
        materials: &mut ResMut<Assets<ColorMaterial>>,
        size: f32,
        color: Color,
        transform: Transform,
    ) -> Ball {
        commands.spawn(MaterialMesh2dBundle {
            mesh: meshes.add(shape::Circle::new(size).into()).into(),
            material: materials.add(ColorMaterial::from(color)),
            transform: transform,
            ..default()
        });

        return Ball {
            transform: transform,
        };
    }
}
