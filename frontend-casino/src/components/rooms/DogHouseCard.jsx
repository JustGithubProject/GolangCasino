import React from 'react';

const DogHouseCard = () => {
    let defaultMatrix = [[10, 7, 6, 4, 1], [4, 6, 5, 4, 6], [3, 1, 3, 3, 4]]
    const [gameBoard, setGameBoard] = useState(defaultMatrix);
    const [username, setUsername] = useState(null);
    const [balance, setBalance] = useState(null);
    const [isSpinning, setIsSpinning] = useState(false);
    const [bet, setBet] = useState(10);
    const [isMusicPlaying, setIsMusicPlaying] = useState(true);
    const audioRef = useRef(null); 

    useEffect(() => {
        const token = localStorage.getItem('token');
        if (token) {
        const decodedToken = jwtDecodeModule.jwtDecode(token);
        const username = decodedToken.username;
        setUsername(username);
        fetchUserBalance(username);
        }

    // Воспроизведение музыки при монтировании компонента
    if (audioRef.current) {
      audioRef.current.loop = true; // Зацикливаем музыку
      audioRef.current.volume = 0.5; // Устанавливаем громкость
      if (isMusicPlaying) {
        audioRef.current.play();
      } else {
        audioRef.current.pause();
      }
    }
    }, [isMusicPlaying]);

    const fetchUserBalance = async (username) => {
    try {
        const response = await fetchWithAuth(`http://localhost:8081/user/name/${username}`);
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        const data = await response.json();
        setBalance(data.Balance);
        } catch (error) {
        console.error('Error fetching user balance:', error);
        }
    };

    const handleSpin = async () => {
        if (bet <= 0) {
        alert('Ставка должна быть больше нуля');
        return;
        }

        setIsSpinning(true);
        setGameBoard(defaultMatrix);

        setTimeout(async () => {
        try {
            const url = `http://localhost:8081/spin-slot-v1/?spinBet=${bet}`;
            const response = await fetchWithAuth(url, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({}),
            });

            if (!response.ok) {
            throw new Error('Network response was not ok');
            }

            let data = await response.json();
            
            // Getting a string from json representing a string of numbers
            let sRow1 = data.sRow1;
            let sRow2 = data.sRow2;
            let sRow3 = data.sRow3;

            // ["2", "3", "4", ... "5"]
            let strRow1Array = sRow1.split(", ");
            let strRow2Array = sRow2.split(", ");
            let strRow3Array = sRow3.split(", ");

            let numRow1Array = strRow1Array.map(Number);
            let numRow2Array = strRow2Array.map(Number);
            let numRow3Array = strRow3Array.map(Number);

            let matrixArray = [
                numRow1Array,
                numRow2Array,
                numRow3Array,
            ];


            setGameBoard(matrixArray);
            setBalance(data.balance);
        } catch (error) {
            console.error('Error fetching spin data:', error);
        } finally {
            setIsSpinning(false);
        }
        }, 500);
    };

    const toggleMusic = () => {
        setIsMusicPlaying(!isMusicPlaying);
    };

    return (
        <div>

        </div>
    );
}