use bevy::{color::palettes::basic::PURPLE, prelude::*};

#[derive(Component)]
struct Player1;

#[derive(Component)]
pub struct Direction {
    pub x: f32,
    pub y: f32,
}

fn main() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_systems(Startup, setup)
        .add_systems(Update, update)
        .add_systems(Update, handle_events)
        .run();
}

fn setup(
    mut commands: Commands,
    mut meshes: ResMut<Assets<Mesh>>,
    mut materials: ResMut<Assets<ColorMaterial>>,
) {
    commands.spawn(Camera2d);

    // Ball
    commands.spawn((
        Mesh2d(meshes.add(Circle::default())),
        MeshMaterial2d(materials.add(Color::from(PURPLE))),
        Transform::default().with_scale(Vec3::splat(15.)),
    ));

    // Player 1
    commands
        .spawn((
            Mesh2d(meshes.add(Rectangle::default())),
            MeshMaterial2d(materials.add(Color::from(PURPLE))),
            Transform::from_xyz(-500., 0., 0.)
                .with_scale(Vec3::new(15., 120., 0.)),
        ))
        .insert(Player1)
        .insert(Direction { x: 0., y: 0. });

    // Player 2
    commands.spawn((
        Mesh2d(meshes.add(Rectangle::default())),
        MeshMaterial2d(materials.add(Color::from(PURPLE))),
        Transform::from_xyz(500., 0., 0.).with_scale(Vec3::new(15., 120., 0.)),
    ));
}

fn update(
    time: Res<Time>,
    mut query: Query<(&Direction, &mut Transform), With<Player1>>,
) {
    for (direction, mut transform) in query.iter_mut() {
        let translation = &mut transform.translation;
        translation.y += direction.y * time.delta_secs() * 200.;
    }
}

fn handle_events(
    kb: Res<ButtonInput<KeyCode>>,
    mut query: Query<&mut Direction, With<Player1>>,
) {
    if let Ok(mut direction) = query.single_mut() {
        direction.y = if kb.pressed(KeyCode::KeyW) {
            1.
        } else if kb.pressed(KeyCode::KeyS) {
            -1.
        } else {
            0.
        }
    }
}
