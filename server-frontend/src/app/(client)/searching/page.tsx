"use client";
import { useState, useEffect } from 'react';
import { Moon, Sun, ArrowLeft, Users, Settings, Clock, MessageCircle, Shield, UserPlus } from 'lucide-react';

export default function MatchmakingPage() {
  const [darkMode, setDarkMode] = useState(true);
  const [matchmakingState, setMatchmakingState] = useState('searching'); // searching, found, ready
  const [timeElapsed, setTimeElapsed] = useState(0);
  const [playerCount, setPlayerCount] = useState(1);
  const [showSettings, setShowSettings] = useState(false);
  
  // Options de matchmaking
  const [gameOptions, setGameOptions] = useState({
    playerCount: '8-10',
    gameMode: 'classique',
    roleBalance: 'équilibré',
    allowSpectators: true
  });

  // Simuler la recherche de joueurs
  useEffect(() => {
    let interval: string | number | NodeJS.Timeout | undefined;
    
    if (matchmakingState === 'searching') {
      interval = setInterval(() => {
        setTimeElapsed(prev => prev + 1);
        
        // Simuler l'arrivée de nouveaux joueurs
        if (playerCount < 10 && Math.random() > 0.7) {
          setPlayerCount(prev => Math.min(prev + 1, 10));
        }
        
        // Simuler que la partie est trouvée après ~20 secondes
        if (timeElapsed > 15 && playerCount >= 8) {
          setMatchmakingState('found');
          setTimeout(() => {
            setMatchmakingState('ready');
          }, 3000);
        }
      }, 1000);
    }
    
    return () => clearInterval(interval);
  }, [matchmakingState, timeElapsed, playerCount]);
  
  const toggleDarkMode = () => {
    setDarkMode(!darkMode);
  };
  
  // Formater le temps écoulé
  const formatTime = (seconds: number) => {
    const mins = Math.floor(seconds / 60);
    const secs = seconds % 60;
    return `${mins}:${secs < 10 ? '0' : ''}${secs}`;
  };
  
  // Redémarrer la recherche
  const restartSearch = () => {
    setTimeElapsed(0);
    setPlayerCount(1);
    setMatchmakingState('searching');
  };
  
  return (
    <div className={`min-h-screen ${darkMode ? 'bg-gray-900 text-gray-100' : 'bg-gray-100 text-gray-900'}`}>
      {/* Fond décoratif */}
      <div className="fixed top-0 left-0 w-full h-full overflow-hidden z-0">
        <div className={`absolute top-0 left-0 w-full h-full ${darkMode ? 'opacity-20' : 'opacity-10'} bg-cover bg-center`}>
          <img src="/api/placeholder/1920/1080" alt="background" className="w-full h-full object-cover" />
        </div>
        <div className={`absolute top-0 left-0 w-full h-full ${darkMode ? 'bg-gradient-to-br from-purple-900/50 to-indigo-900/50' : 'bg-gradient-to-br from-amber-100/30 to-orange-100/30'}`}></div>
      </div>

      {/* Header */}
      <header className={`sticky top-0 z-50 ${darkMode ? 'bg-gray-900/80 backdrop-blur-sm' : 'bg-white/80 backdrop-blur-sm'} shadow-md`}>
        <div className="container mx-auto px-4 py-3 flex items-center justify-between">
          <div className="flex items-center space-x-3">
            <button className={`p-2 rounded-full ${darkMode ? 'hover:bg-gray-800' : 'hover:bg-gray-200'}`}>
              <ArrowLeft size={24} />
            </button>
            <h1 className={`text-xl font-bold ${darkMode ? 'text-indigo-300' : 'text-amber-800'}`}>Recherche de partie</h1>
          </div>

          <div className="flex items-center space-x-3">
            <button 
              onClick={() => setShowSettings(!showSettings)}
              className={`p-2 rounded-full ${darkMode ? 'hover:bg-gray-800' : 'hover:bg-gray-200'} ${showSettings ? darkMode ? 'bg-gray-800' : 'bg-gray-200' : ''}`}
            >
              <Settings size={20} />
            </button>
            <button 
              onClick={toggleDarkMode}
              className={`p-2 rounded-full ${darkMode ? 'bg-gray-800 text-yellow-400 hover:bg-gray-700' : 'bg-gray-200 text-indigo-600 hover:bg-gray-300'}`}
            >
              {darkMode ? <Sun size={20} /> : <Moon size={20} />}
            </button>
          </div>
        </div>
      </header>

      {/* Contenu principal */}
      <main className="container mx-auto px-4 py-8 relative z-10 flex flex-col items-center justify-center">
        {/* Options de matchmaking (affichées conditionnellement) */}
        {showSettings && (
          <div className={`w-full max-w-lg mb-8 rounded-lg shadow-lg ${darkMode ? 'bg-gray-800/90' : 'bg-white/90'} backdrop-blur-sm p-6`}>
            <h2 className="text-xl font-semibold mb-4">Options de partie</h2>
            
            <div className="space-y-4">
              <div>
                <label className="block mb-2 text-sm font-medium">Nombre de joueurs</label>
                <select 
                  value={gameOptions.playerCount}
                  onChange={(e) => setGameOptions({...gameOptions, playerCount: e.target.value})}
                  className={`w-full p-2 rounded-md ${
                    darkMode 
                      ? 'bg-gray-700 border-gray-600 text-white' 
                      : 'bg-gray-50 border-gray-300 text-gray-900'
                  }`}
                >
                  <option value="6-8">Petit groupe (6-8)</option>
                  <option value="8-10">Groupe moyen (8-10)</option>
                  <option value="12-16">Grand groupe (12-16)</option>
                </select>
              </div>
              
              <div>
                <label className="block mb-2 text-sm font-medium">Mode de jeu</label>
                <select 
                  value={gameOptions.gameMode}
                  onChange={(e) => setGameOptions({...gameOptions, gameMode: e.target.value})}
                  className={`w-full p-2 rounded-md ${
                    darkMode 
                      ? 'bg-gray-700 border-gray-600 text-white' 
                      : 'bg-gray-50 border-gray-300 text-gray-900'
                  }`}
                >
                  <option value="classique">Classique</option>
                  <option value="rapide">Mode rapide</option>
                  <option value="roles_avances">Rôles avancés</option>
                </select>
              </div>
              
              <div>
                <label className="block mb-2 text-sm font-medium">Équilibre des rôles</label>
                <select 
                  value={gameOptions.roleBalance}
                  onChange={(e) => setGameOptions({...gameOptions, roleBalance: e.target.value})}
                  className={`w-full p-2 rounded-md ${
                    darkMode 
                      ? 'bg-gray-700 border-gray-600 text-white' 
                      : 'bg-gray-50 border-gray-300 text-gray-900'
                  }`}
                >
                  <option value="équilibré">Équilibré</option>
                  <option value="pro_village">Favorable au village</option>
                  <option value="pro_loups">Favorable aux loups</option>
                  <option value="aléatoire">Totalement aléatoire</option>
                </select>
              </div>
              
              <div className="flex items-center">
                <input
                  id="spectators"
                  type="checkbox"
                  checked={gameOptions.allowSpectators}
                  onChange={() => setGameOptions({...gameOptions, allowSpectators: !gameOptions.allowSpectators})}
                  className={`h-4 w-4 rounded ${
                    darkMode 
                      ? 'bg-gray-700 border-gray-600 text-indigo-500' 
                      : 'bg-gray-50 border-gray-300 text-amber-600'
                  }`}
                />
                <label htmlFor="spectators" className="ml-2 text-sm">
                  Autoriser les spectateurs
                </label>
              </div>
            </div>
            
            <div className="mt-6 flex justify-end">
              <button 
                onClick={() => setShowSettings(false)}
                className={`py-2 px-4 rounded-md ${
                  darkMode ? 'bg-indigo-600 hover:bg-indigo-700' : 'bg-amber-600 hover:bg-amber-700'
                } text-white`}
              >
                Appliquer
              </button>
            </div>
          </div>
        )}

        {/* Zone principale de matchmaking */}
        <div className={`w-full max-w-lg rounded-lg shadow-lg ${darkMode ? 'bg-gray-800/90' : 'bg-white/90'} backdrop-blur-sm p-6 mb-6`}>
          <div className="text-center mb-8">
            <h2 className="text-2xl font-bold mb-2">
              {matchmakingState === 'searching' && 'Recherche de joueurs...'}
              {matchmakingState === 'found' && 'Partie trouvée !'}
              {matchmakingState === 'ready' && 'Prêt à commencer !'}
            </h2>
            
            <p className={`text-sm ${darkMode ? 'text-gray-400' : 'text-gray-600'}`}>
              <Clock size={16} className="inline mr-1" /> Temps écoulé: {formatTime(timeElapsed)}
            </p>
          </div>

          {/* Animation différente selon l'état */}
          <div className="flex justify-center mb-8">
            {matchmakingState === 'searching' && (
              <div className="relative w-64 h-64">
                {/* Cercle extérieur animé */}
                <div className="absolute inset-0 rounded-full border-4 border-transparent border-t-indigo-500 border-r-indigo-500 animate-spin"></div>
                
                {/* Cercle central avec lune */}
                <div className={`absolute inset-6 rounded-full flex items-center justify-center ${darkMode ? 'bg-gray-700' : 'bg-gray-100'}`}>
                  <div className="relative">
                    <Moon size={48} className={`${darkMode ? 'text-yellow-300' : 'text-amber-500'}`} />
                    
                    {/* Étoiles animées autour de la lune */}
                    <div className="absolute top-0 right-0 h-2 w-2 bg-yellow-200 rounded-full animate-pulse"></div>
                    <div className="absolute bottom-2 left-0 h-1 w-1 bg-yellow-200 rounded-full animate-pulse" style={{animationDelay: '0.5s'}}></div>
                    <div className="absolute top-6 left-0 h-2 w-2 bg-yellow-200 rounded-full animate-pulse" style={{animationDelay: '1s'}}></div>
                  </div>
                </div>
                
                {/* Nombre de joueurs positionnés autour du cercle */}
                {Array.from({ length: 10 }).map((_, index) => {
                  const angle = (index * 36) * Math.PI / 180;
                  const x = Math.cos(angle) * 100 + 100;
                  const y = Math.sin(angle) * 100 + 100;
                  
                  return (
                    <div 
                      key={index} 
                      className={`absolute w-10 h-10 rounded-full flex items-center justify-center transition-all duration-500 ${
                        index < playerCount 
                          ? darkMode ? 'bg-indigo-700 text-white' : 'bg-amber-200 text-amber-800'
                          : darkMode ? 'bg-gray-700 text-gray-500' : 'bg-gray-200 text-gray-400'
                      }`}
                      style={{
                        transform: `translate(${x}px, ${y}px) translate(-20px, -20px)`,
                        opacity: index < playerCount ? 1 : 0.5
                      }}
                    >
                      <Users size={16} />
                    </div>
                  );
                })}
              </div>
            )}
            
            {matchmakingState === 'found' && (
              <div className="w-64 h-64 flex items-center justify-center">
                <div className={`w-48 h-48 rounded-full flex items-center justify-center ${darkMode ? 'bg-green-900/50' : 'bg-green-100'} animate-pulse`}>
                  <div className={`w-40 h-40 rounded-full flex items-center justify-center ${darkMode ? 'bg-green-800/50' : 'bg-green-200'}`}>
                    <Users size={64} className={`${darkMode ? 'text-green-300' : 'text-green-600'}`} />
                  </div>
                </div>
              </div>
            )}
            
            {matchmakingState === 'ready' && (
              <div className="w-64 h-64 flex items-center justify-center">
                <div className={`w-48 h-48 rounded-full flex items-center justify-center ${darkMode ? 'bg-green-900/50' : 'bg-green-100'}`}>
                  <div className={`w-40 h-40 rounded-full flex items-center justify-center ${darkMode ? 'bg-green-800/50' : 'bg-green-200'} animate-bounce`}>
                    <Shield size={64} className={`${darkMode ? 'text-green-300' : 'text-green-600'}`} />
                  </div>
                </div>
              </div>
            )}
          </div>

          {/* Information sur les joueurs */}
          <div className={`text-center mb-8 ${matchmakingState === 'searching' ? 'block' : 'hidden'}`}>
            <p className="text-lg font-medium">
              <Users size={20} className="inline mr-2" /> 
              {playerCount}/10 joueurs en attente
            </p>
            <p className={`text-sm mt-2 ${darkMode ? 'text-gray-400' : 'text-gray-600'}`}>
              Recherche d'une partie {gameOptions.gameMode} avec {gameOptions.playerCount} joueurs...
            </p>
          </div>

          {/* Message informatif selon l'état */}
          {matchmakingState === 'found' && (
            <div className="text-center bg-green-800/20 border border-green-700/30 rounded-lg p-4 mb-6">
              <p className="font-medium text-green-500">Une partie a été trouvée !</p>
              <p className={`text-sm mt-1 ${darkMode ? 'text-gray-300' : 'text-gray-600'}`}>
                Préparation de la partie en cours...
              </p>
            </div>
          )}
          
          {matchmakingState === 'ready' && (
            <div className="text-center bg-green-800/20 border border-green-700/30 rounded-lg p-4 mb-6">
              <p className="font-medium text-green-500">Partie prête à commencer !</p>
              <p className={`text-sm mt-1 ${darkMode ? 'text-gray-300' : 'text-gray-600'}`}>
                Mode de jeu: {gameOptions.gameMode} • {playerCount} joueurs
              </p>
            </div>
          )}

          {/* Bouton d'action */}
          {matchmakingState === 'searching' && (
            <button className={`w-full py-3 rounded-md font-medium ${
              darkMode ? 'bg-red-600 hover:bg-red-700' : 'bg-red-500 hover:bg-red-600'
            } text-white`}>
              Annuler la recherche
            </button>
          )}
          
          {matchmakingState === 'found' && (
            <div className="flex space-x-4">
              <button className={`flex-1 py-3 rounded-md font-medium ${
                darkMode ? 'bg-gray-700 hover:bg-gray-600' : 'bg-gray-300 hover:bg-gray-400'
              } text-white`}>
                Annuler
              </button>
              <button className={`flex-1 py-3 rounded-md font-medium animate-pulse ${
                darkMode ? 'bg-green-600 hover:bg-green-700' : 'bg-green-500 hover:bg-green-600'
              } text-white`}>
                Se préparer
              </button>
            </div>
          )}
          
          {matchmakingState === 'ready' && (
            <button className={`w-full py-3 rounded-md font-medium ${
              darkMode ? 'bg-green-600 hover:bg-green-700' : 'bg-green-500 hover:bg-green-600'
            } text-white`}>
              Entrer dans la partie
            </button>
          )}
        </div>

        {/* Suggestions d'action pendant l'attente */}
        {matchmakingState === 'searching' && (
          <div className={`w-full max-w-lg rounded-lg shadow-lg ${darkMode ? 'bg-gray-800/90' : 'bg-white/90'} backdrop-blur-sm p-4`}>
            <h3 className="text-lg font-medium mb-3">Pendant l'attente...</h3>
            
            <div className="grid grid-cols-1 sm:grid-cols-2 gap-3">
              <button className={`flex items-center p-3 rounded-md ${
                darkMode ? 'bg-gray-700 hover:bg-gray-600' : 'bg-gray-200 hover:bg-gray-300'
              }`}>
                <MessageCircle size={20} className="mr-3" />
                <span>Ouvrir le chat</span>
              </button>
              
              <button className={`flex items-center p-3 rounded-md ${
                darkMode ? 'bg-gray-700 hover:bg-gray-600' : 'bg-gray-200 hover:bg-gray-300'
              }`}>
                <UserPlus size={20} className="mr-3" />
                <span>Inviter des amis</span>
              </button>
            </div>
          </div>
        )}
      </main>
    </div>
  );
}