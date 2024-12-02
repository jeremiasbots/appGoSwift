//
//  ContentView.swift
//  AppGo
//
//  Created by devep on 30/11/24.
//

import SwiftUI
import GoSwift
import CoreData

struct ContentView: View {
    @Environment(\.managedObjectContext) private var viewContext
    
    var body: some View {
        VStack {
            AsyncImage(url: URL(string: GoSwiftTheNewMovieImage())) { phase in
                switch phase {
                case .failure:
                    ProgressView()
                case .success(let image):
                    image
                        .resizable()
                default:
                    ProgressView()
                }
            }
                .frame(width: 350, height: 450)
                .clipShape(.rect(cornerRadius: 25))
            Text(GoSwiftTheNewMovie())
                .font(.title2)
                .padding()
            Text("La película se estrena en \(GoSwiftTheNewMovieDays()) días")
                .font(.headline)
                .padding()
        }
    }
}


#Preview {
    ContentView().environment(\.managedObjectContext, PersistenceController.preview.container.viewContext)
}
