"use client";
import { useState, useEffect } from 'react';
import { MessageSquare, Moon, Sun, Users, X } from 'lucide-react';

// Types pour notre jeu
type Role = 'villageois' | 'loup-garou' | 'voyante' | 'sorcière' | 'chasseur';
type Phase = 'jour' | 'nuit' | 'vote';
type Message = {
  id: number;
  sender: string;
  content: string;
  timestamp: Date;
  isSystemMessage: boolean;
};

// Component pour la carte du joueur
const PlayerCard = ({ role, isRevealed }: { role: Role, isRevealed: boolean }) => {
  const [flipped, setFlipped] = useState(!isRevealed);
  
  const getRoleImage = (role: Role) => {
    // Dans une application réelle, vous utiliseriez de vraies images
    switch(role) {
      case 'loup-garou': return '/api/placeholder/160/220';
      case 'voyante': return '/api/placeholder/160/220';
      case 'sorcière': return '/api/placeholder/160/220';
      case 'chasseur': return '/api/placeholder/160/220';
      default: return '/api/placeholder/160/220';
    }
  };

  const getRoleColor = (role: Role) => {
    switch(role) {
      case 'loup-garou': return 'bg-red-700';
      case 'voyante': return 'bg-purple-700';
      case 'sorcière': return 'bg-green-700';
      case 'chasseur': return 'bg-yellow-700';
      default: return 'bg-blue-700';
    }
  };

  return (
    <div 
      className="relative cursor-pointer w-40 h-56 perspective-1000"
      onClick={() => setFlipped(!flipped)}
    >
      <div className={`absolute w-full h-full transition-transform duration-500 ${flipped ? 'rotate-y-180' : ''}`}>
        {/* Face avant (dos de la carte) */}
        <div className={`absolute w-full h-full backface-hidden bg-indigo-800 rounded-lg flex items-center justify-center ${flipped ? 'hidden' : ''}`}>
          <div className="text-white text-lg font-bold">Votre Rôle</div>
        </div>
        
        {/* Face arrière (rôle du joueur) */}
        <div className={`absolute w-full h-full backface-hidden rotate-y-180 rounded-lg overflow-hidden ${getRoleColor(role)} ${!flipped ? 'hidden' : ''}`}>
          <div className="absolute inset-0 flex flex-col items-center justify-between p-2">
            <div className="text-white text-sm font-bold uppercase">{role}</div>
            <div className="w-32 h-32 bg-white rounded-full overflow-hidden flex items-center justify-center">
              <img src={getRoleImage(role)} alt={role} className="w-full h-full object-cover" />
            </div>
            <div className="text-white text-xs">Cliquez pour retourner</div>
          </div>
        </div>
      </div>
    </div>
  );
};

// Component pour le chat
const GameChat = ({ messages, sendMessage }: { messages: Message[], sendMessage: (content: string) => void }) => {
  const [messageInput, setMessageInput] = useState('');
  const [isChatOpen, setIsChatOpen] = useState(true);

  const handleSendMessage = () => {
    if (messageInput.trim()) {
      sendMessage(messageInput);
      setMessageInput('');
    }
  };

  const handleKeyPress = (e: { key: string; }) => {
    if (e.key === 'Enter') {
      handleSendMessage();
    }
  };

  return (
    <div className={`absolute bottom-0 right-0 w-80 transition-all duration-300 ${isChatOpen ? 'h-96' : 'h-12'} bg-gray-800 bg-opacity-90 rounded-tl-lg shadow-lg flex flex-col`}>
      {/* En-tête du chat */}
      <div 
        className="flex items-center justify-between px-4 py-2 bg-gray-900 rounded-tl-lg cursor-pointer"
        onClick={() => setIsChatOpen(!isChatOpen)}
      >
        <div className="flex items-center space-x-2">
          <MessageSquare size={18} className="text-white" />
          <span className="text-white font-medium">Chat du Village</span>
        </div>
        <button className="text-gray-400 hover:text-white">
          {isChatOpen ? <X size={18} /> : <span className="text-xs">+</span>}
        </button>
      </div>
      
      {/* Corps du chat */}
      {isChatOpen && (
        <>
          <div className="flex-1 overflow-y-auto p-3 space-y-2">
            {messages.map((msg) => (
              <div key={msg.id} className={`${msg.isSystemMessage ? 'bg-red-900 bg-opacity-50' : 'bg-gray-700'} p-2 rounded-lg max-w-[90%] ${msg.sender === 'Vous' ? 'ml-auto' : ''}`}>
                <div className="flex justify-between items-baseline mb-1">
                  <span className={`font-bold text-xs ${msg.isSystemMessage ? 'text-red-300' : 'text-blue-300'}`}>
                    {msg.sender}
                  </span>
                  <span className="text-gray-400 text-xs">
                    {new Date(msg.timestamp).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
                  </span>
                </div>
                <p className="text-white text-sm">{msg.content}</p>
              </div>
            ))}
          </div>
          
          {/* Saisie de message */}
          <div className="p-2 bg-gray-900 flex">
            <input
              type="text"
              value={messageInput}
              onChange={(e) => setMessageInput(e.target.value)}
              onKeyPress={handleKeyPress}
              className="flex-1 bg-gray-700 text-white rounded-l px-3 py-2 focus:outline-none"
              placeholder="Écrivez un message..."
            />
            <button 
              onClick={handleSendMessage} 
              className="bg-indigo-600 hover:bg-indigo-700 text-white px-4 rounded-r"
            >
              Envoyer
            </button>
          </div>
        </>
      )}
    </div>
  );
};

// Component principal pour le village
const VillageMap = ({ phase, players, currentDay }: { phase: Phase, players: Array<{ id: number, name: string, isAlive: boolean }>, currentDay: number }) => {
  return (
    <div className="relative w-full h-full overflow-hidden">
      {/* Fond du village */}
      <div className="absolute inset-0 bg-gradient-to-b from-indigo-900 to-gray-900">
        {/* Effet jour/nuit */}
        <div className={`absolute inset-0 transition-opacity duration-3000 ${phase === 'jour' ? 'opacity-0' : 'opacity-80'} bg-gray-900`}></div>
        
        {/* Ciel et effet astronomique */}
        <div className="absolute top-0 left-0 right-0 h-32">
          {phase === 'jour' ? (
            <div className="absolute top-8 right-16 w-20 h-20 bg-yellow-300 rounded-full blur-md animate-pulse"></div>
          ) : (
            <>
              <div className="absolute top-12 left-1/4 w-16 h-16 bg-gray-100 rounded-full blur-sm"></div>
              <div className="stars absolute inset-0"></div>
            </>
          )}
        </div>
        
        {/* Silhouettes de maisons */}
        <div className="absolute bottom-0 left-0 right-0 h-64">
          <div className="absolute bottom-0 left-10 w-24 h-32 bg-gray-800"></div>
          <div className="absolute bottom-0 left-10 w-24 h-6 bg-gray-900 transform rotate-45 origin-bottom-left"></div>
          <div className="absolute bottom-0 left-40 w-28 h-40 bg-gray-800"></div>
          <div className="absolute bottom-0 left-40 w-28 h-8 bg-gray-900 transform -rotate-12 origin-bottom-left"></div>
          <div className="absolute bottom-0 right-20 w-32 h-36 bg-gray-800"></div>
          <div className="absolute bottom-0 right-20 w-32 h-8 bg-gray-900 transform rotate-12 origin-bottom-right"></div>
          <div className="absolute bottom-0 right-60 w-24 h-28 bg-gray-800"></div>
          <div className="absolute bottom-0 right-60 w-24 h-6 bg-gray-900 transform -rotate-30 origin-bottom-right"></div>
        </div>
        
        {/* Place du village (cercle central) */}
        <div className="absolute left-1/2 bottom-40 transform -translate-x-1/2 w-64 h-64 rounded-full bg-gray-700 bg-opacity-30 flex items-center justify-center">
          {phase === 'vote' && (
            <div className="animate-pulse text-red-500 font-bold text-lg">Vote en cours</div>
          )}
          {phase === 'nuit' && (
            <div className="animate-pulse text-blue-300 font-bold text-lg">Les loups-garous rôdent...</div>
          )}
        </div>
        
        {/* Avatars des joueurs placés en cercle */}
        {players.map((player, index) => {
          const angle = (index / players.length) * Math.PI * 2;
          const x = Math.cos(angle) * 120 + 32;
          const y = Math.sin(angle) * 120 + 32;
          
          return (
            <div 
              key={player.id}
              className={`absolute flex flex-col items-center transition-all duration-500 transform -translate-x-1/2 -translate-y-1/2`}
              style={{ 
                left: `calc(50% + ${x}px)`, 
                bottom: `calc(40% + ${y}px)`,
                opacity: player.isAlive ? 1 : 0.4
              }}
            >
              <div className={`w-12 h-12 rounded-full bg-gray-300 flex items-center justify-center overflow-hidden ${!player.isAlive ? 'grayscale' : ''}`}>
                <img src={`/api/placeholder/48/48`} alt={player.name} className="w-full h-full object-cover" />
              </div>
              <div className={`mt-1 px-2 py-1 rounded bg-gray-800 bg-opacity-70 text-xs ${player.isAlive ? 'text-white' : 'text-gray-400'}`}>
                {player.name}
                {!player.isAlive && " †"}
              </div>
            </div>
          );
        })}
      </div>
    </div>
  );
};

// Component pour l'indicateur de phase
const PhaseIndicator = ({ phase, currentDay }: { phase: Phase, currentDay: number }) => {
  return (
    <div className="absolute top-4 left-1/2 transform -translate-x-1/2 flex items-center space-x-3 bg-gray-800 bg-opacity-80 px-4 py-2 rounded-full">
      <div className="flex items-center space-x-2">
        {phase === 'jour' ? (
          <Sun size={20} className="text-yellow-300" />
        ) : (
          <Moon size={20} className="text-blue-300" />
        )}
        <span className={`font-bold text-lg ${phase === 'jour' ? 'text-yellow-300' : 'text-blue-300'}`}>
          {phase === 'jour' ? 'Jour' : 'Nuit'} {currentDay}
        </span>
      </div>
      
      {phase === 'vote' && (
        <div className="flex items-center space-x-2 text-red-400">
          <Users size={20} />
          <span className="font-bold">Vote</span>
        </div>
      )}
    </div>
  );
};

// Component principal
export default function WerewolfGame() {
  const [playerRole, setPlayerRole] = useState<Role>('loup-garou');
  const [isRoleRevealed, setIsRoleRevealed] = useState(false);
  const [gamePhase, setGamePhase] = useState<Phase>('jour');
  const [currentDay, setCurrentDay] = useState(1);
  const [messages, setMessages] = useState<Message[]>([
    {
      id: 1,
      sender: 'Système',
      content: 'Bienvenue dans le village! La partie va bientôt commencer.',
      timestamp: new Date(),
      isSystemMessage: true
    },
    {
      id: 2,
      sender: 'Système',
      content: 'Regardez votre carte pour découvrir votre rôle secret.',
      timestamp: new Date(),
      isSystemMessage: true
    }
  ]);
  
  const [players, setPlayers] = useState([
    { id: 1, name: 'Vous', isAlive: true },
    { id: 2, name: 'Jean', isAlive: true },
    { id: 3, name: 'Sophie', isAlive: true },
    { id: 4, name: 'Michel', isAlive: true },
    { id: 5, name: 'Lucie', isAlive: false },
    { id: 6, name: 'Thomas', isAlive: true },
    { id: 7, name: 'Émilie', isAlive: true },
    { id: 8, name: 'Antoine', isAlive: true }
  ]);

  // Pour simuler le changement de phase
  useEffect(() => {
    const timer = setTimeout(() => {
      if (gamePhase === 'jour') {
        setGamePhase('vote');
        addSystemMessage('Le village doit maintenant voter pour éliminer un suspect!');
      } else if (gamePhase === 'vote') {
        setGamePhase('nuit');
        addSystemMessage('La nuit tombe sur le village. Les loups-garous se réveillent...');
      } else {
        setGamePhase('jour');
        setCurrentDay(currentDay + 1);
        addSystemMessage(`Le jour ${currentDay + 1} se lève sur le village.`);
      }
    }, 20000); // Changer toutes les 20 secondes pour la démo

    return () => clearTimeout(timer);
  }, [gamePhase, currentDay]);

  // Pour révéler la carte au début
  useEffect(() => {
    const timer = setTimeout(() => {
      setIsRoleRevealed(true);
      addSystemMessage('Votre rôle a été révélé! Ne le montrez à personne.');
    }, 3000);

    return () => clearTimeout(timer);
  }, []);

  const addSystemMessage = (content: string) => {
    const newMessage: Message = {
      id: messages.length + 1,
      sender: 'Système',
      content,
      timestamp: new Date(),
      isSystemMessage: true
    };
    setMessages([...messages, newMessage]);
  };

  const sendMessage = (content: string) => {
    const newMessage: Message = {
      id: messages.length + 1,
      sender: 'Vous',
      content,
      timestamp: new Date(),
      isSystemMessage: false
    };
    setMessages([...messages, newMessage]);
  };

  return (
    <div className="relative w-full h-screen bg-gray-900 overflow-hidden">
      {/* Carte du village (en arrière-plan) */}
      <VillageMap phase={gamePhase} players={players} currentDay={currentDay} />
      
      {/* Indicateur de phase */}
      <PhaseIndicator phase={gamePhase} currentDay={currentDay} />
      
      {/* Carte du joueur */}
      <div className="absolute top-4 left-4">
        <PlayerCard role={playerRole} isRevealed={isRoleRevealed} />
      </div>
      
      {/* Chat */}
      <GameChat messages={messages} sendMessage={sendMessage} />
      
      {/* Styles globaux */}
      <style jsx global>{`
        @keyframes twinkle {
          0% { opacity: 0.3; }
          50% { opacity: 1; }
          100% { opacity: 0.3; }
        }
        
        .stars {
          background-image: radial-gradient(white 1px, transparent 1px);
          background-size: 20px 20px;
        }
        
        .stars::before {
          content: '';
          position: absolute;
          top: 0;
          left: 0;
          right: 0;
          bottom: 0;
          background-image: radial-gradient(white 1px, transparent 1px);
          background-size: 30px 30px;
          background-position: 10px 10px;
          animation: twinkle 4s infinite;
        }
        
        .rotate-y-180 {
          transform: rotateY(180deg);
        }
        
        .perspective-1000 {
          perspective: 1000px;
        }
        
        .backface-hidden {
          backface-visibility: hidden;
        }
        
        .duration-3000 {
          transition-duration: 3000ms;
        }
      `}</style>
    </div>
  );
}