//
//  Recommendation.swift
//  L2Play
//
//  Created by Lukasz Fabia on 11/12/2024.
//

import Foundation

final class RecommendationService {
    var allGames: [Game]
    var userGames: [GameWithState]
    
    init(allGames: [Game], userGames: [GameWithState]) {
        self.allGames = allGames
        self.userGames = userGames
    }
    
    /// Predict games based on players played games
    /// - Returns: games ids
    func predict() -> [String] {
        var userTags = Set<String>()
        var userGameIds = Set<String>()
        
        userGames.forEach { gameState in
            userGameIds.insert(gameState.gameID.uuidString)
            
            if gameState.state != .dropped {
                if let game = allGames.first(where: { $0.id == gameState.gameID }) {
                    userTags.formUnion(game.tags)
                }
            }
        }
        
        var allTags = allGames.flatMap { $0.tags }
        allTags.append(contentsOf: Array(userTags))
        
        let vectors = countVectorizer(tags: allGames.map { $0.tags })

        
        let userVector = vectors.last!
        let gameVectors = vectors.dropLast()
    
        let cosineSim = gameVectors.map { cosineSimilarity(vec1: userVector, vec2: $0) }
        
        
        let sortedIndices = cosineSim.enumerated().sorted { $0.element > $1.element }.map { $0.offset }
        
        var recommendedGames = [String]()
        
        for idx in sortedIndices {
            if recommendedGames.count >= 3 {
                break
            }
            
            guard idx < allGames.count else { continue }
            
            let gameId = allGames[idx].id.uuidString
            if !userGameIds.contains(gameId) {
                recommendedGames.append(gameId)
            }
        }
        
        return recommendedGames
    }
    
    private func countVectorizer(tags: [[String]]) -> [[Int]] {
        var tagSet = Set<String>()
        
        for tagList in tags {
            for tag in tagList {
                tagSet.insert(tag)
            }
        }
        
        let tagArray = Array(tagSet)
        var vectors = [[Int]]()
        
        for tagList in tags {
            var vector = [Int](repeating: 0, count: tagArray.count)
            for tag in tagList {
                if let index = tagArray.firstIndex(of: tag) {
                    vector[index] += 1
                }
            }
            vectors.append(vector)
        }
        
        return vectors
    }
    
    private func cosineSimilarity(vec1: [Int], vec2: [Int]) -> Double {
        let dotProduct = zip(vec1, vec2).map { $0.0 * $0.1 }.reduce(0, +)
        let magnitude1 = sqrt(vec1.map { Double($0 * $0) }.reduce(0, +))
        let magnitude2 = sqrt(vec2.map { Double($0 * $0) }.reduce(0, +))
        
        return Double(dotProduct) / (magnitude1 * magnitude2)
    }
}