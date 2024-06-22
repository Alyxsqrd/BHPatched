# Brick Hill Patched

This script allows you to play the new Brick Hill Workshop made with Godot Engine offline.

[Click here](https://github.com/Alyxsqrd/WTBArchiver/releases/latest) to download the latest version.

## Modifications

-   [x] Default scene is updated to `workshop.tscn`, probably the easiest patch ever, lol.

-   [x] Skybox assets are manually copied during startup, due to a race condition in the BH source code.

-   [x] Completely removed the `bin` folder, this script makes those spooky binaries useless.

-   [ ] Normally, I would've made a built-in HTTP proxy for this sort of project... but since the official APIs are still online (as of now) and the game is 69% unfinished, this won't be necessary until the inevitable shutdown happens.

## Credits

| Name                | Details                               |
| :------------------ | :------------------------------------ |
| **Alyx**            | Reverse engineering and patching      |
| **Enderspearl184**  | Sharing latest unmodified beta build  |
| **dargy**           | OAuth access and URI scheme handling  |
| **StarManTheGamer** | Pre-release testing and feedback      |
| **spacebuilder**    | Barnes and Noble incident :trollface: |
