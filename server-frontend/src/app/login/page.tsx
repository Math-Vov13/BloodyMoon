"use client";
import { useState } from 'react';
import { Moon, Sun, Eye, EyeOff } from 'lucide-react';

export default function LoginPage() {
  const [darkMode, setDarkMode] = useState(true);
  const [showPassword, setShowPassword] = useState(false);
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [rememberMe, setRememberMe] = useState(false);

  const toggleDarkMode = () => {
    setDarkMode(!darkMode);
  };

  const handleLogin = () => {
    // Ici serait la logique d'authentification
    console.log('Tentative de connexion avec:', { email, password, rememberMe });
  };

  return (
    <div className={`flex min-h-screen flex-col items-center justify-center p-4 ${darkMode ? 'bg-gray-900 text-gray-100' : 'bg-gray-100 text-gray-900'}`}>
      {/* Fond décoratif */}
      <div className="absolute top-0 left-0 w-full h-full overflow-hidden z-0">
        <div className={`absolute top-0 left-0 w-full h-full ${darkMode ? 'opacity-20' : 'opacity-10'} bg-cover bg-center`}>
          <img src="/api/placeholder/1920/1080" alt="background" className="w-full h-full object-cover" />
        </div>
        <div className={`absolute top-0 left-0 w-full h-full ${darkMode ? 'bg-gradient-to-br from-purple-900/50 to-indigo-900/50' : 'bg-gradient-to-br from-amber-100/30 to-orange-100/30'}`}></div>
      </div>
      
      {/* Toggle mode jour/nuit */}
      <button 
        onClick={toggleDarkMode}
        className={`absolute top-4 right-4 p-2 rounded-full z-10 ${darkMode ? 'bg-gray-800 text-yellow-400 hover:bg-gray-700' : 'bg-gray-200 text-indigo-600 hover:bg-gray-300'}`}
      >
        {darkMode ? <Sun size={20} /> : <Moon size={20} />}
      </button>

      {/* Logo et titre */}
      <div className="relative z-10 mb-8 flex flex-col items-center">
        <div className={`w-24 h-24 rounded-full flex items-center justify-center mb-4 ${darkMode ? 'bg-indigo-900' : 'bg-amber-100'}`}>
          <div className={`w-20 h-20 rounded-full flex items-center justify-center ${darkMode ? 'bg-indigo-800' : 'bg-amber-200'}`}>
            <Moon size={36} className={darkMode ? 'text-yellow-300' : 'text-amber-700'} />
          </div>
        </div>
        <h1 className={`text-3xl font-bold ${darkMode ? 'text-indigo-300' : 'text-amber-800'}`}>Loup-Garou</h1>
        <p className={`mt-1 ${darkMode ? 'text-gray-400' : 'text-gray-600'}`}>La nuit tombe sur le village...</p>
      </div>

      {/* Formulaire de connexion */}
      <div className={`w-full max-w-md relative z-10 p-8 rounded-lg shadow-lg ${darkMode ? 'bg-gray-800/80 backdrop-blur-sm' : 'bg-white/90 backdrop-blur-sm'}`}>
        <h2 className={`text-xl font-semibold mb-6 text-center ${darkMode ? 'text-gray-200' : 'text-gray-800'}`}>
          Connexion
        </h2>
        
        <div className="space-y-6">
          <div>
            <label htmlFor="email" className={`block mb-2 text-sm font-medium ${darkMode ? 'text-gray-300' : 'text-gray-700'}`}>
              Email
            </label>
            <input
              id="email"
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className={`w-full p-3 rounded-md border ${
                darkMode 
                  ? 'bg-gray-700 border-gray-600 text-white placeholder-gray-400 focus:border-indigo-500' 
                  : 'bg-gray-50 border-gray-300 text-gray-900 placeholder-gray-500 focus:border-amber-500'
              } focus:outline-none focus:ring-2 ${
                darkMode ? 'focus:ring-indigo-500/50' : 'focus:ring-amber-500/50'
              }`}
              placeholder="votre@email.com"
            />
          </div>
          
          <div>
            <label htmlFor="password" className={`block mb-2 text-sm font-medium ${darkMode ? 'text-gray-300' : 'text-gray-700'}`}>
              Mot de passe
            </label>
            <div className="relative">
              <input
                id="password"
                type={showPassword ? 'text' : 'password'}
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                className={`w-full p-3 rounded-md border ${
                  darkMode 
                    ? 'bg-gray-700 border-gray-600 text-white placeholder-gray-400 focus:border-indigo-500' 
                    : 'bg-gray-50 border-gray-300 text-gray-900 placeholder-gray-500 focus:border-amber-500'
                } focus:outline-none focus:ring-2 ${
                  darkMode ? 'focus:ring-indigo-500/50' : 'focus:ring-amber-500/50'
                }`}
                placeholder="••••••••"
              />
              <button 
                type="button"
                className={`absolute right-3 top-1/2 transform -translate-y-1/2 ${darkMode ? 'text-gray-400 hover:text-gray-300' : 'text-gray-500 hover:text-gray-700'}`}
                onClick={() => setShowPassword(!showPassword)}
              >
                {showPassword ? <EyeOff size={20} /> : <Eye size={20} />}
              </button>
            </div>
          </div>
          
          <div className="flex items-center justify-between">
            <div className="flex items-center">
              <input
                id="remember"
                type="checkbox"
                checked={rememberMe}
                onChange={() => setRememberMe(!rememberMe)}
                className={`h-4 w-4 rounded ${
                  darkMode 
                    ? 'bg-gray-700 border-gray-600 text-indigo-500 focus:ring-indigo-600' 
                    : 'bg-gray-50 border-gray-300 text-amber-600 focus:ring-amber-500'
                }`}
              />
              <label htmlFor="remember" className={`ml-2 text-sm ${darkMode ? 'text-gray-300' : 'text-gray-700'}`}>
                Se souvenir de moi
              </label>
            </div>
            <a href="#" className={`text-sm font-medium ${darkMode ? 'text-indigo-400 hover:text-indigo-300' : 'text-amber-600 hover:text-amber-700'}`}>
              Mot de passe oublié?
            </a>
          </div>
          
          <button
            onClick={handleLogin}
            className={`w-full py-3 px-4 rounded-md font-medium transition duration-200 ${
              darkMode 
                ? 'bg-indigo-600 hover:bg-indigo-700 text-white' 
                : 'bg-amber-600 hover:bg-amber-700 text-white'
            }`}
          >
            Se connecter
          </button>
        </div>
        
        <div className="mt-6 text-center">
          <p className={darkMode ? 'text-gray-400' : 'text-gray-600'}>
            Pas encore de compte?{' '}
            <a href="#" className={`font-medium ${darkMode ? 'text-indigo-400 hover:text-indigo-300' : 'text-amber-600 hover:text-amber-700'}`}>
              S'inscrire
            </a>
          </p>
        </div>
      </div>
      
      <footer className="relative z-10 mt-8 text-center text-sm">
        <p className={darkMode ? 'text-gray-500' : 'text-gray-600'}>
          © 2025 Loup-Garou Online | <a href="#" className={`${darkMode ? 'text-gray-400 hover:text-gray-300' : 'text-gray-600 hover:text-gray-800'}`}>Mentions légales</a>
        </p>
      </footer>
    </div>
  );
}