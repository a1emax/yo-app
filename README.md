# Yo App

**[Youngine](https://github.com/a1emax/youngine) App Template**

## Compilation

### Windows

Run `make build-windows`.

You can add the same commands for other desktop platforms, but for them:
* C compiler is required;
* cross-compilation is not an option.

### Android

Install Android SDK/NDK (for the first time).

Run `make install-ebitenmobile` (for the first time).

Run `make build-android`.

Open `app/android` project in Android Studio and run it on emulator or your device.

## Layout

Order of packages reflects possible dependencies between them - lower packages may depend on upper ones,
but not vice versa.

* **res** - embedded file system containing resources (assets, configs, etc).
* **pkg** - imported packages.
    * **domain** - domain logic.
    * **global** - global entities.
        * **consts** - arbitrary constants.
        * **vars** - arbitrary variables.
        * **tools** - Youngine tools, logger, RNG, etc.
        * **assets** - static assets.
    * **window** - GUI.
    * **kernel** - control kernel.
* **cmd** - compiled service packages (if any).
* **app** - compiled application packages.
    * **desktop** - main for Windows, Linux and macOS.
    * **android_intern** - library for Android (compiled to AAR).
    * **android** - Android Studio project.
