.PHONY: install-ebitenmobile
install-ebitenmobile:
	go install github.com/hajimehoshi/ebiten/v2/cmd/ebitenmobile@v2.8.6

.PHONY: build-windows
build-windows:
	mkdir -p .local
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build \
		-o .local/yo-app-windows.exe \
		yo-app/app/desktop

.PHONY: build-android
build-android:
	mkdir -p .local
	ebitenmobile bind \
		-target android \
		-androidapi 24 \
		-javapkg com.github.a1emax.yoapp.go \
		-o .local/yo-app-android-intern.aar \
		yo-app/app/android_intern
	cp .local/yo-app-android-intern.aar app/android/intern/default.aar

.PHONY: grab-android
grab-android:
	cp app/android/app/build/outputs/apk/debug/app-debug.apk .local/yo-app-android.apk
