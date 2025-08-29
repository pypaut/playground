use bevy::{color::palettes::basic::PURPLE, prelude::*, app::AppExit};
use bevy_rapier2d::prelude::*;

const BALL_RADIUS: f32 = 15.;
const BALL_SPEED: f32 = 300.;
const PLAYER_SIZE: f32 = 120.;
const PLAYER_SPEED: f32 = 500.;

#[derive(Component)]
struct Player;

#[derive(Component)]
struct Player1;

#[derive(Component)]
struct Player2;

#[derive(Component)]
struct Ball;

#[derive(Component)]
pub struct Direction {
    pub x: f32,
    pub y: f32,
}

#[derive(Component)]
struct Buttons {
    up: KeyCode,
    down: KeyCode,
}

fn main() {
    App::new()
        .add_plugins(DefaultPlugins)
        .add_plugins(RapierDebugRenderPlugin::default())
        .add_plugins(RapierPhysicsPlugin::<NoUserData>::pixels_per_meter(100.0))
        .add_systems(Startup, setup)
        // .add_systems(Update, print_ball_altitude)
        .run();
}

fn setup(
    mut commands: Commands,
    mut meshes: ResMut<Assets<Mesh>>,
    mut materials: ResMut<Assets<ColorMaterial>>,
) {
    // Add a camera so we can see the debug-render.
    commands.spawn(Camera2d::default());

    // Ball
    commands
        .spawn(RigidBody::Dynamic)
        .insert(Collider::ball(5.0))
        .insert(Restitution::coefficient(1.0))
        .insert(Transform::from_xyz(0.0, 0.0, 0.0))
        .insert(Ball)
        .insert(Direction { x: 0., y: 0. })
        .insert(GravityScale(1.0))
        .insert(Mesh2d(meshes.add(Circle::default())))
        .insert(MeshMaterial2d(materials.add(Color::from(PURPLE))));

    // Bottom wall
    commands
        .spawn(Collider::cuboid(1000.0, 50.0))
        .insert(Transform::from_xyz(0.0, -400.0, 0.0));

    // Top wall
    commands
        .spawn(Collider::cuboid(1000.0, 50.0))
        .insert(Transform::from_xyz(0.0, 400.0, 0.0));

    // Left wall
    commands
        .spawn(Collider::cuboid(50.0, 1000.0))
        .insert(Transform::from_xyz(-500.0, 0.0, 0.0));

    // Right wall
    commands
        .spawn(Collider::cuboid(50.0, 1000.0))
        .insert(Transform::from_xyz(500.0, 0.0, 0.0));

    // Player 1
    // commands
    //     .spawn(Collider::capsule(Vect{x: 0., y: 50.}, Vect{x: 0., y: -50.}, 10.0))
    //     .insert(Transform::from_xyz(500.0, 0.0, 0.0))
    //     .insert(Player)
    //     .insert(Player1)
    //     .insert(Direction { x: 0., y: 0. })
    //     .insert(Buttons {
    //         up: KeyCode::KeyW,
    //         down: KeyCode::KeyS,
    //     });
}

fn ball_update(
    mut query: Query<(&mut Transform, &mut Restitution), With<Ball>>,
) {
    for (mut transform, mut restitution) in query.iter_mut() {

    }
}

// fn print_ball_altitude(positions: Query<&Transform, With<RigidBody>>) {
//     for transform in positions.iter() {
//         println!("Ball altitude: {}", transform.translation.y);
//     }
// }
