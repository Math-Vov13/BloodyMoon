"use client";
import { SetStateAction, useState } from 'react';
import { Moon, Sun, Users, Award, Book, Play, LogIn, Menu, X, ChevronDown, ChevronUp } from 'lucide-react';

export default function HomePage() {
  const [darkMode, setDarkMode] = useState(true);
  const [menuOpen, setMenuOpen] = useState(false);
  const [openSection, setOpenSection] = useState('');

  const toggleDarkMode = () => {
    setDarkMode(!darkMode);
  };

  const toggleSection = (section: SetStateAction<string>) => {
    setOpenSection(openSection === section ? '' : section);
  };

  const roles = [
    { 
      id: 'villageois', 
      name: 'Villageois', 
      description: 'Simple habitant sans pouvoir particulier. Son objectif est de démasquer et éliminer les loups-garous.',
      team: 'village',
      color: darkMode ? 'bg-blue-700' : 'bg-blue-500'
    },
    { 
      id: 'loup', 
      name: 'Loup-Garou', 
      description: 'Se réveille la nuit pour dévorer un villageois. Doit se faire passer pour un villageois le jour.',
      team: 'loups',
      color: darkMode ? 'bg-red-800' : 'bg-red-600'
    },
    { 
      id: 'voyante', 
      name: 'Voyante', 
      description: 'Peut découvrir la véritable identité d\'un joueur chaque nuit.',
      team: 'village',
      color: darkMode ? 'bg-purple-700' : 'bg-purple-500'
    },
    { 
      id: 'sorciere', 
      name: 'Sorcière', 
      description: 'Possède deux potions : une pour guérir, une pour tuer.',
      team: 'village',
      color: darkMode ? 'bg-green-700' : 'bg-green-500'
    },
    { 
      id: 'chasseur', 
      name: 'Chasseur', 
      description: 'Lorsqu\'il meurt, il peut immédiatement éliminer un autre joueur.',
      team: 'village',
      color: darkMode ? 'bg-orange-700' : 'bg-orange-500'
    },
    { 
      id: 'cupidon', 
      name: 'Cupidon', 
      description: 'Désigne deux amoureux qui partageront le même destin.',
      team: 'village',
      color: darkMode ? 'bg-pink-700' : 'bg-pink-500'
    }
  ];

  return (
    <div className={`min-h-screen ${darkMode ? 'bg-gray-900 text-gray-100' : 'bg-gray-100 text-gray-900'}`}>
      {/* Fond décoratif */}
      <div className="fixed top-0 left-0 w-full h-full overflow-hidden z-0">
        <div className={`absolute top-0 left-0 w-full h-full ${darkMode ? 'opacity-20' : 'opacity-10'} bg-cover bg-center`}>
          <img src="/api/placeholder/1920/1080" alt="background" className="w-full h-full object-cover" />
        </div>
        <div className={`absolute top-0 left-0 w-full h-full ${darkMode ? 'bg-gradient-to-br from-purple-900/50 to-indigo-900/50' : 'bg-gradient-to-br from-amber-100/30 to-orange-100/30'}`}></div>
      </div>

      {/* Navigation */}
      <header className={`sticky top-0 z-50 ${darkMode ? 'bg-gray-900/80 backdrop-blur-sm' : 'bg-white/80 backdrop-blur-sm'} shadow-md`}>
        <div className="container mx-auto px-4 py-3 flex items-center justify-between">
          <div className="flex items-center space-x-3">
            <div className={`h-10 w-10 rounded-full flex items-center justify-center ${darkMode ? 'bg-indigo-800' : 'bg-amber-200'}`}>
              <Moon size={24} className={darkMode ? 'text-yellow-300' : 'text-amber-700'} />
            </div>
            <h1 className={`text-xl font-bold ${darkMode ? 'text-indigo-300' : 'text-amber-800'}`}>Loup-Garou</h1>
          </div>

          {/* Desktop Navigation */}
          <nav className="hidden md:flex items-center space-x-6">
            <a href="#regles" className={`font-medium hover:underline ${darkMode ? 'text-gray-200 hover:text-white' : 'text-gray-700 hover:text-black'}`}>Règles</a>
            <a href="#roles" className={`font-medium hover:underline ${darkMode ? 'text-gray-200 hover:text-white' : 'text-gray-700 hover:text-black'}`}>Rôles</a>
            <a href="#classement" className={`font-medium hover:underline ${darkMode ? 'text-gray-200 hover:text-white' : 'text-gray-700 hover:text-black'}`}>Classement</a>
            <a href="#apropos" className={`font-medium hover:underline ${darkMode ? 'text-gray-200 hover:text-white' : 'text-gray-700 hover:text-black'}`}>À propos</a>
          </nav>

          <div className="flex items-center space-x-3">
            {/* Toggle mode jour/nuit */}
            <button 
              onClick={toggleDarkMode}
              className={`p-2 rounded-full ${darkMode ? 'bg-gray-800 text-yellow-400 hover:bg-gray-700' : 'bg-gray-200 text-indigo-600 hover:bg-gray-300'}`}
            >
              {darkMode ? <Sun size={20} /> : <Moon size={20} />}
            </button>

            {/* Boutons d'action */}
            <div className="hidden md:flex space-x-3">
              <button className={`py-2 px-4 rounded-md font-medium ${darkMode ? 'bg-indigo-600 hover:bg-indigo-700 text-white' : 'bg-amber-600 hover:bg-amber-700 text-white'}`}>
                <LogIn size={16} className="inline mr-1" /> Se connecter
              </button>
              <button className={`py-2 px-4 rounded-md font-medium ${darkMode ? 'bg-red-600 hover:bg-red-700 text-white' : 'bg-red-600 hover:bg-red-700 text-white'}`}>
                <Play size={16} className="inline mr-1" /> Jouer
              </button>
            </div>

            {/* Mobile menu button */}
            <button 
              onClick={() => setMenuOpen(!menuOpen)}
              className="md:hidden p-2 rounded-md"
            >
              {menuOpen ? <X size={24} /> : <Menu size={24} />}
            </button>
          </div>
        </div>

        {/* Mobile menu */}
        {menuOpen && (
          <div className={`md:hidden ${darkMode ? 'bg-gray-800' : 'bg-white'} py-4 px-6 space-y-4`}>
            <a href="#regles" className="block py-2 font-medium" onClick={() => setMenuOpen(false)}>Règles</a>
            <a href="#roles" className="block py-2 font-medium" onClick={() => setMenuOpen(false)}>Rôles</a>
            <a href="#classement" className="block py-2 font-medium" onClick={() => setMenuOpen(false)}>Classement</a>
            <a href="#apropos" className="block py-2 font-medium" onClick={() => setMenuOpen(false)}>À propos</a>
            
            <div className="flex flex-col space-y-3 pt-4 border-t border-gray-700">
              <button className={`py-2 px-4 rounded-md font-medium ${darkMode ? 'bg-indigo-600 hover:bg-indigo-700 text-white' : 'bg-amber-600 hover:bg-amber-700 text-white'}`}>
                <LogIn size={16} className="inline mr-2" /> Se connecter
              </button>
              <button className={`py-2 px-4 rounded-md font-medium ${darkMode ? 'bg-red-600 hover:bg-red-700 text-white' : 'bg-red-600 hover:bg-red-700 text-white'}`}>
                <Play size={16} className="inline mr-2" /> Jouer
              </button>
            </div>
          </div>
        )}
      </header>

      {/* Contenu principal */}
      <main className="container mx-auto px-4 py-12 relative z-10">
        {/* Hero section */}
        <section className="text-center mb-16">
          <div className="max-w-3xl mx-auto">
            <h1 className={`text-4xl md:text-5xl font-bold mb-6 ${darkMode ? 'text-indigo-300' : 'text-amber-800'}`}>
              Bienvenue au village des Loups-Garous
            </h1>
            <p className="text-xl mb-8">
              Démasquez les loups-garous ou dévorez les villageois dans ce jeu classique de bluff et de déduction sociale.
            </p>
            <div className="flex flex-col sm:flex-row justify-center gap-4">
              <button className={`py-3 px-8 rounded-md text-lg font-medium ${darkMode ? 'bg-indigo-600 hover:bg-indigo-700 text-white' : 'bg-amber-600 hover:bg-amber-700 text-white'}`}>
                <LogIn size={20} className="inline mr-2" /> Se connecter
              </button>
              <button className={`py-3 px-8 rounded-md text-lg font-medium ${darkMode ? 'bg-red-600 hover:bg-red-700 text-white' : 'bg-red-600 hover:bg-red-700 text-white'}`}>
                <Play size={20} className="inline mr-2" /> Jouer maintenant
              </button>
            </div>
          </div>
        </section>

        {/* Cartes de rôle */}
        <section id="roles" className="mb-16">
          <h2 className={`text-3xl font-bold mb-8 text-center ${darkMode ? 'text-gray-200' : 'text-gray-800'}`}>
            <Users size={24} className="inline mr-2" /> Découvrez les rôles
          </h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
            {roles.map(role => (
              <div 
                key={role.id}
                className={`rounded-lg overflow-hidden shadow-lg transition-transform transform hover:scale-105 ${darkMode ? 'bg-gray-800' : 'bg-white'}`}
              >
                <div className={`h-20 ${role.color} flex items-center justify-center`}>
                  <h3 className="text-2xl font-bold text-white">{role.name}</h3>
                </div>
                <div className="p-6">
                  <div className={`text-sm font-semibold mb-3 inline-block px-3 py-1 rounded-full ${
                    role.team === 'village' 
                      ? darkMode ? 'bg-blue-900/50 text-blue-200' : 'bg-blue-100 text-blue-800' 
                      : darkMode ? 'bg-red-900/50 text-red-200' : 'bg-red-100 text-red-800'
                  }`}>
                    {role.team === 'village' ? 'Camp du Village' : 'Camp des Loups'}
                  </div>
                  <p className={`${darkMode ? 'text-gray-300' : 'text-gray-600'}`}>
                    {role.description}
                  </p>
                </div>
              </div>
            ))}
          </div>
          <div className="mt-6 text-center">
            <button className={`inline-flex items-center py-2 px-4 rounded-md font-medium ${darkMode ? 'bg-gray-700 hover:bg-gray-600 text-gray-200' : 'bg-gray-200 hover:bg-gray-300 text-gray-800'}`}>
              Voir tous les rôles <ChevronDown size={16} className="ml-2" />
            </button>
          </div>
        </section>

        {/* Règles du jeu */}
        <section id="regles" className="mb-16">
          <h2 className={`text-3xl font-bold mb-8 text-center ${darkMode ? 'text-gray-200' : 'text-gray-800'}`}>
            <Book size={24} className="inline mr-2" /> Règles du jeu
          </h2>
          
          <div className={`rounded-lg shadow-lg ${darkMode ? 'bg-gray-800/80' : 'bg-white/90'} backdrop-blur-sm`}>
            {/* Section Présentation */}
            <div className="border-b border-gray-700">
              <button 
                onClick={() => toggleSection('presentation')}
                className={`w-full flex justify-between items-center p-5 text-left font-medium ${darkMode ? 'hover:bg-gray-700/50' : 'hover:bg-gray-100'}`}
              >
                <span className="text-xl">Présentation du jeu</span>
                {openSection === 'presentation' ? <ChevronUp size={20} /> : <ChevronDown size={20} />}
              </button>
              
              {openSection === 'presentation' && (
                <div className="p-5 pt-0">
                  <p className="mb-4">
                    Loup-Garou est un jeu de société dans lequel chaque joueur incarne un villageois ou un loup-garou. 
                    Le but des villageois est de découvrir qui sont les loups-garous parmi eux et de les éliminer. 
                    Les loups-garous, quant à eux, cherchent à dévorer tous les villageois sans être démasqués.
                  </p>
                  <p>
                    Le jeu alterne entre deux phases : la nuit, où les loups-garous et certains villageois avec des pouvoirs 
                    spéciaux agissent secrètement, et le jour, où tous les joueurs débattent pour décider qui éliminer par vote.
                  </p>
                </div>
              )}
            </div>
            
            {/* Section Déroulement */}
            <div className="border-b border-gray-700">
              <button 
                onClick={() => toggleSection('deroulement')}
                className={`w-full flex justify-between items-center p-5 text-left font-medium ${darkMode ? 'hover:bg-gray-700/50' : 'hover:bg-gray-100'}`}
              >
                <span className="text-xl">Déroulement d'une partie</span>
                {openSection === 'deroulement' ? <ChevronUp size={20} /> : <ChevronDown size={20} />}
              </button>
              
              {openSection === 'deroulement' && (
                <div className="p-5 pt-0">
                  <ol className="list-decimal list-inside space-y-3">
                    <li><strong>Distribution des rôles</strong> : Chaque joueur reçoit secrètement un rôle.</li>
                    <li><strong>Phase de nuit</strong> : Les loups-garous choisissent une victime. Certains rôles spéciaux utilisent leurs pouvoirs.</li>
                    <li><strong>Phase de jour</strong> : Les joueurs découvrent qui a été éliminé pendant la nuit, puis débattent et votent pour éliminer un suspect.</li>
                    <li><strong>Fin de la partie</strong> : Le jeu continue jusqu'à ce que tous les loups-garous soient éliminés (victoire des villageois) ou que les loups-garous soient aussi nombreux que les villageois (victoire des loups-garous).</li>
                  </ol>
                </div>
              )}
            </div>
            
            {/* Section Conseils */}
            <div>
              <button 
                onClick={() => toggleSection('conseils')}
                className={`w-full flex justify-between items-center p-5 text-left font-medium ${darkMode ? 'hover:bg-gray-700/50' : 'hover:bg-gray-100'}`}
              >
                <span className="text-xl">Conseils et stratégies</span>
                {openSection === 'conseils' ? <ChevronUp size={20} /> : <ChevronDown size={20} />}
              </button>
              
              {openSection === 'conseils' && (
                <div className="p-5 pt-0">
                  <ul className="space-y-3">
                    <li><strong>Pour les villageois</strong> : Observez attentivement le comportement de chacun. Prenez des notes sur qui accuse qui et comment les joueurs réagissent.</li>
                    <li><strong>Pour les loups-garous</strong> : Essayez de vous fondre parmi les villageois. Accusez habilement pour détourner les soupçons.</li>
                    <li><strong>Pour tous</strong> : Adaptez votre stratégie selon votre rôle et l'évolution de la partie. La communication est clé !</li>
                  </ul>
                </div>
              )}
            </div>
          </div>
        </section>

        {/* Classement */}
        <section id="classement" className="mb-16">
          <h2 className={`text-3xl font-bold mb-8 text-center ${darkMode ? 'text-gray-200' : 'text-gray-800'}`}>
            <Award size={24} className="inline mr-2" /> Classement des joueurs
          </h2>
          
          <div className={`rounded-lg shadow-lg overflow-hidden ${darkMode ? 'bg-gray-800/80' : 'bg-white/90'} backdrop-blur-sm`}>
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead className={`${darkMode ? 'bg-gray-700' : 'bg-gray-100'}`}>
                  <tr>
                    <th className="px-6 py-3 text-left">Rang</th>
                    <th className="px-6 py-3 text-left">Joueur</th>
                    <th className="px-6 py-3 text-left">Parties</th>
                    <th className="px-6 py-3 text-left">Victoires</th>
                    <th className="px-6 py-3 text-left">Ratio</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-700">
                  {[
                    { rank: 1, name: "Loup_Solitaire", games: 245, wins: 187, ratio: "76%" },
                    { rank: 2, name: "VoyanteSupreme", games: 198, wins: 142, ratio: "72%" },
                    { rank: 3, name: "ChasseurFou", games: 312, wins: 218, ratio: "70%" },
                    { rank: 4, name: "SorciereDu78", games: 176, wins: 121, ratio: "69%" },
                    { rank: 5, name: "CupidonFlèche", games: 220, wins: 141, ratio: "64%" }
                  ].map((player, idx) => (
                    <tr key={idx} className={`${darkMode ? 'hover:bg-gray-700' : 'hover:bg-gray-50'}`}>
                      <td className="px-6 py-4">{player.rank}</td>
                      <td className="px-6 py-4 font-medium">{player.name}</td>
                      <td className="px-6 py-4">{player.games}</td>
                      <td className="px-6 py-4">{player.wins}</td>
                      <td className="px-6 py-4">{player.ratio}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
          
          <div className="mt-6 text-center">
            <button className={`inline-flex items-center py-2 px-4 rounded-md font-medium ${darkMode ? 'bg-gray-700 hover:bg-gray-600 text-gray-200' : 'bg-gray-200 hover:bg-gray-300 text-gray-800'}`}>
              Voir le classement complet <ChevronDown size={16} className="ml-2" />
            </button>
          </div>
        </section>

        {/* Fonctionnalités */}
        <section className="mb-16">
          <h2 className={`text-3xl font-bold mb-8 text-center ${darkMode ? 'text-gray-200' : 'text-gray-800'}`}>
            Fonctionnalités principales
          </h2>
          
          <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
            <div className={`p-6 rounded-lg shadow-lg ${darkMode ? 'bg-gray-800/80' : 'bg-white/90'} backdrop-blur-sm`}>
              <div className={`mb-4 inline-flex p-3 rounded-full ${darkMode ? 'bg-indigo-900/50' : 'bg-indigo-100'}`}>
                <Users size={24} className={darkMode ? 'text-indigo-300' : 'text-indigo-600'} />
              </div>
              <h3 className={`text-xl font-bold mb-2 ${darkMode ? 'text-gray-200' : 'text-gray-800'}`}>Multijoueur en temps réel</h3>
              <p className={darkMode ? 'text-gray-300' : 'text-gray-600'}>
                Jouez avec vos amis ou avec des joueurs du monde entier dans des parties dynamiques et immersives.
              </p>
            </div>
            
            <div className={`p-6 rounded-lg shadow-lg ${darkMode ? 'bg-gray-800/80' : 'bg-white/90'} backdrop-blur-sm`}>
              <div className={`mb-4 inline-flex p-3 rounded-full ${darkMode ? 'bg-green-900/50' : 'bg-green-100'}`}>
                <Award size={24} className={darkMode ? 'text-green-300' : 'text-green-600'} />
              </div>
              <h3 className={`text-xl font-bold mb-2 ${darkMode ? 'text-gray-200' : 'text-gray-800'}`}>Système de progression</h3>
              <p className={darkMode ? 'text-gray-300' : 'text-gray-600'}>
                Gagnez de l'expérience, débloquez des avatars et grimpez dans le classement mondial.
              </p>
            </div>
            
            <div className={`p-6 rounded-lg shadow-lg ${darkMode ? 'bg-gray-800/80' : 'bg-white/90'} backdrop-blur-sm`}>
              <div className={`mb-4 inline-flex p-3 rounded-full ${darkMode ? 'bg-amber-900/50' : 'bg-amber-100'}`}>
                <Moon size={24} className={darkMode ? 'text-amber-300' : 'text-amber-600'} />
              </div>
              <h3 className={`text-xl font-bold mb-2 ${darkMode ? 'text-gray-200' : 'text-gray-800'}`}>Modes de jeu variés</h3>
              <p className={darkMode ? 'text-gray-300' : 'text-gray-600'}>
                Découvrez différentes variantes du jeu avec des rôles exclusifs et des règles spéciales.
              </p>
            </div>
          </div>
        </section>
        
        {/* Call to action */}
        <section className={`text-center p-10 rounded-lg shadow-lg mb-16 ${darkMode ? 'bg-indigo-900/60' : 'bg-amber-100/80'} backdrop-blur-sm`}>
          <h2 className={`text-3xl font-bold mb-4 ${darkMode ? 'text-indigo-200' : 'text-amber-800'}`}>
            Prêt à rejoindre la meute ?
          </h2>
          <p className="text-lg mb-8 max-w-2xl mx-auto">
            Créez votre compte gratuitement et plongez dans l'univers mystérieux du Loup-Garou. 
            Des centaines de joueurs vous attendent pour des parties palpitantes !
          </p>
          <div className="flex flex-col sm:flex-row justify-center gap-4">
            <button className={`py-3 px-8 rounded-md text-lg font-medium ${darkMode ? 'bg-gray-800 hover:bg-gray-700 text-white' : 'bg-white hover:bg-gray-100 text-gray-800'}`}>
              En savoir plus
            </button>
            <button className={`py-3 px-8 rounded-md text-lg font-medium ${darkMode ? 'bg-red-600 hover:bg-red-700 text-white' : 'bg-red-600 hover:bg-red-700 text-white'}`}>
              <Play size={20} className="inline mr-2" /> Jouer maintenant
            </button>
          </div>
        </section>
      </main>

      {/* Footer */}
      <footer className={`py-8 relative z-10 ${darkMode ? 'bg-gray-800/90' : 'bg-gray-100/90'} backdrop-blur-sm`}>
        <div className="container mx-auto px-4">
          <div className="flex flex-col md:flex-row justify-between items-center">
            <div className="flex items-center space-x-3 mb-4 md:mb-0">
              <div className={`h-8 w-8 rounded-full flex items-center justify-center ${darkMode ? 'bg-indigo-800' : 'bg-amber-200'}`}>
                <Moon size={18} className={darkMode ? 'text-yellow-300' : 'text-amber-700'} />
              </div>
              <span className={`text-lg font-bold ${darkMode ? 'text-indigo-300' : 'text-amber-800'}`}>Loup-Garou</span>
            </div>
            
            <div className="flex flex-wrap justify-center gap-x-8 gap-y-2 mb-4 md:mb-0">
              <a href="#" className={`hover:underline ${darkMode ? 'text-gray-300' : 'text-gray-600'}`}>Règles</a>
              <a href="#" className={`hover:underline ${darkMode ? 'text-gray-300' : 'text-gray-600'}`}>Rôles</a>
              <a href="#" className={`hover:underline ${darkMode ? 'text-gray-300' : 'text-gray-600'}`}>Classement</a>
              <a href="#" className={`hover:underline ${darkMode ? 'text-gray-300' : 'text-gray-600'}`}>Aide</a>
              <a href="#" className={`hover:underline ${darkMode ? 'text-gray-300' : 'text-gray-600'}`}>Contact</a>
            </div>
            
            <div className={`text-sm ${darkMode ? 'text-gray-400' : 'text-gray-500'}`}>
              © 2025 Loup-Garou Online. Tous droits réservés.
            </div>
          </div>
        </div>
      </footer>
    </div>
  );
}