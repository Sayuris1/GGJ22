components {
  id: "player"
  component: "/player/player.script"
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
  properties {
    id: "type"
    value: "black"
    type: PROPERTY_TYPE_HASH
  }
}
components {
  id: "change_l"
  component: "/player/black_change.particlefx"
  position {
    x: -4.0
    y: -21.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.57357645
    w: 0.81915206
  }
}
components {
  id: "change_r"
  component: "/player/black_change.particlefx"
  position {
    x: -4.0
    y: -21.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: -0.57357645
    w: 0.81915206
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "tile_set: \"/assets/black.tilesource\"\n"
  "default_animation: \"idle\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "blend_mode: BLEND_MODE_ALPHA\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "collision_shape: \"\"\n"
  "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"black_player\"\n"
  "mask: \"black_hit\"\n"
  "mask: \"black_win\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "    }\n"
  "    rotation {\n"
  "      x: 0.0\n"
  "      y: 0.0\n"
  "      z: 0.0\n"
  "      w: 1.0\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 11.5\n"
  "  data: 19.0\n"
  "  data: 10.0\n"
  "}\n"
  "linear_damping: 0.0\n"
  "angular_damping: 0.0\n"
  "locked_rotation: false\n"
  "bullet: false\n"
  ""
  position {
    x: 0.0
    y: 0.0
    z: 0.0
  }
  rotation {
    x: 0.0
    y: 0.0
    z: 0.0
    w: 1.0
  }
}
