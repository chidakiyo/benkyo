// swift-tools-version:4.2
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription

let package = Package(
    name: "Server",
    dependencies: [
        .package(url: "https://github.com/IBM-Swift/Kitura", from: "2.6.0")
    ],
    targets: [
        .target(
            name: "Server",
            dependencies: ["Kitura"])
    ]
)