plugins {
    kotlin("jvm")
}

group = "org.okapi.server"
version = "1.0.0"

dependencies {
    // Align versions of all Kotlin components
    implementation(platform("org.jetbrains.kotlin:kotlin-bom"))
    // Use the Kotlin JDK 8 standard library.
    implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")
    // okhttp client dependency
    implementation("org.okapi.lib:data-driven-okhttp:1.0.1")
}