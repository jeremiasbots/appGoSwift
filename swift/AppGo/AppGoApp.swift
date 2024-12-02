//
//  AppGoApp.swift
//  AppGo
//
//  Created by devep on 30/11/24.
//

import SwiftUI

@main
struct AppGoApp: App {
    let persistenceController = PersistenceController.shared

    var body: some Scene {
        WindowGroup {
            ContentView()
                .environment(\.managedObjectContext, persistenceController.container.viewContext)
        }
    }
}
