import React, { useState, useEffect, useRef } from 'react';
import styled, { css, keyframes } from 'styled-components'; 
import backgroundImage from '../../images/MainBackground.jpg';
import doghousebackground from '../../images/slots-backgrounds/dog-house-background.png';
import doghouseGameBoardBackground from '../../images/slots-backgrounds/dog-house-gameboard-background.png';
import backgroundMusic from '../../images/doghouse-items/music/doghouse-music.mp3';
import Header from '../header/Header';
import * as jwtDecodeModule from 'jwt-decode';
import { fetchWithAuth } from '../auth_components/fetchWrapper';

import boxerImage from '../../images/doghouse-items/boxer.png';
import aceImage from '../../images/doghouse-items/ace.png';
import kingImage from '../../images/doghouse-items/King.png';
import boneImage from '../../images/doghouse-items/bone.png';
import queenImage from '../../images/doghouse-items/queen.png';
import collarImage from '../../images/doghouse-items/collar.png'
import jackImage from '../../images/doghouse-items/jack.png';
import tenImage from '../../images/doghouse-items/ten.png';
import dachshundImage from '../../images/doghouse-items/dachshund.png';
import pugImage from '../../images/doghouse-items/pug.png';
import pawsImage from '../../images/doghouse-items/paws.png';
import spitzImage from '../../images/doghouse-items/spitz.png';

// Стили для Wrapper
const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  background: linear-gradient(to bottom, rgba(255, 64, 108, 0.85), rgba(255, 75, 43, 0.85));
  background-image: url(${backgroundImage});
  background-size: cover;
  background-position: center;
  background-attachment: fixed;
  height: 100vh;
  width: 100vw;
  box-sizing: border-box;
  padding: 60px 20px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.7);
  overflow: hidden;
  backdrop-filter: blur(4px);
`;

// Анимация для падения символов
const fallAnimation = keyframes`
  0% {
    transform: translateY(-300%) rotate(0deg) scale(0.7);
    opacity: 0;
    box-shadow: 0px 0px 0px rgba(0, 0, 0, 0.3);
  }
  40% {
    transform: translateY(50%) rotate(90deg) scale(1.1);
    opacity: 0.6;
    box-shadow: 0px 10px 20px rgba(0, 0, 0, 0.4);
  }
  70% {
    transform: translateY(10%) rotate(180deg) scale(1.05);
    opacity: 0.8;
    box-shadow: 0px 15px 25px rgba(0, 0, 0, 0.5);
  }
  85% {
    transform: translateY(0) rotate(270deg) scale(1);
    opacity: 0.95;
    box-shadow: 0px 20px 30px rgba(0, 0, 0, 0.6);
  }
  100% {
    transform: translateY(0) rotate(360deg) scale(1);
    opacity: 1;
    box-shadow: 0px 25px 35px rgba(0, 0, 0, 0.7);
  }
`;

// Анимация для перемещения GameBoard вниз
const boardAnimation = keyframes`
  0% { transform: translateY(-150%); }
  50% { transform: translateY(10%); }
  70% { transform: translateY(-3%); }
  90% { transform: translateY(3%); }
  100% { transform: translateY(0); }
`;

// Стили для GameBoard
const GameBoard = styled.div`
  display: grid;
  grid-template-columns: repeat(5, 1fr); 
  grid-template-rows: repeat(3, 1fr); 
  grid-gap: 20px; 
  background-image: url(${doghouseGameBoardBackground});
  padding: 25px; 
  border-radius: 20px; 
  margin-top: 40px; 
  box-shadow: 0 15px 35px rgba(0, 0, 0, 0.6);
  border: 2px solid rgba(102, 102, 102, 0.8);
  position: relative;
  overflow: hidden;
  animation: ${props => props.isSpinning ? css`${boardAnimation} 1.5s cubic-bezier(0.52, 0.04, 0.37, 1) both` : 'none'};
`;

const bounceAnimation = keyframes`
  0% { transform: translateY(0); }
  50% { transform: translateY(-12px); }
  100% { transform: translateY(0); }
`;

// Стили для символов
const Symbol = styled.div`
  width: 90px; 
  height: 90px; 
  display: flex;
  justify-content: center;
  align-items: center;
  background: transparent;
  border-radius: 18px; 
  font-size: 34px; 
  color: #fff;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.5);
  animation: ${bounceAnimation} 1s infinite, ${props => props.isSpinning ? css`${fallAnimation} 1s cubic-bezier(0.52, 0.04, 0.37, 1) both` : 'none'};
  animation-delay: ${props => props.delay || '0s'};
`;

// Стили для кнопок
const Button = styled.button`
  margin-top: 25px;
  padding: 14px 28px;
  background: #28a745;
  color: #fff;
  border: 2px solid rgba(68, 68, 68, 0.8);
  border-radius: 10px;
  cursor: pointer;
  font-size: 22px;
  font-weight: bold;
  transition: background 0.3s ease, transform 0.3s ease, box-shadow 0.3s ease;

  &:hover {
    background: #218838;
    transform: scale(1.1);
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.5);
  }
`;

// Стили для текста баланса
const BalanceText = styled.p`
  margin-top: 20px;
  font-size: 24px;
  color: #ffdf00;
  background: rgba(0, 0, 0, 0.85);
  padding: 14px 28px;
  border-radius: 10px;
  border: 2px solid rgba(85, 85, 85, 0.8);
  text-shadow: 2px 2px 6px rgba(0, 0, 0, 0.9);
`;

// Стили для заголовка
const Title = styled.h1`
  color: #fff;
  margin-top: 30px;
  font-size: 42px;
  text-shadow: 3px 3px 6px rgba(0, 0, 0, 0.9), 0 0 18px #ff4500;
  border: 2px solid rgba(68, 68, 68, 0.8);
  padding: 18px;
  border-radius: 10px;
  background: rgba(0, 0, 0, 0.75);
`;

// Стили для внутреннего контейнера
const InnerWrapper = styled.div`
  border: 2px solid rgba(68, 68, 68, 0.8);
  border-radius: 18px;
  padding: 25px;
  background-image: url(${doghousebackground});
  background-size: cover;
  background-position: center;
  backdrop-filter: blur(8px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.6);
`;

// Стили для инпута ставки
const BetInput = styled.input`
  margin-top: 20px;
  padding: 14px;
  font-size: 22px;
  border: 2px solid rgba(85, 85, 85, 0.8);
  border-radius: 10px;
  background: rgba(0, 0, 0, 0.85);
  color: #fff;
  width: 100px;
  text-align: center;
`;

const MusicButton = styled.button`
  margin-top: 25px;
  padding: 14px 28px;
  background: #007bff;
  color: #fff;
  border: 2px solid rgba(68, 68, 68, 0.8);
  border-radius: 10px;
  cursor: pointer;
  font-size: 22px;
  font-weight: bold;
  transition: background 0.3s ease, transform 0.3s ease, box-shadow 0.3s ease;

  &:hover {
    background: #0056b3;
    transform: scale(1.1);
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.5);
  }
`;


const symbols = [
    { id: 1, name: 'Ten', image: tenImage },
    { id: 2, name: 'Jack', image: jackImage },
    { id: 3, name: 'Queen', image: queenImage },
    { id: 4, name: 'King', image:  kingImage},
    { id: 5, name: 'Ace', image: aceImage },
    { id: 6, name: 'Bone', image: boneImage },
    { id: 7, name: 'Collar', image: collarImage },
    { id: 8, name: 'Dachshund', image: dachshundImage },
    { id: 9, name: 'Pug', image: pugImage },
    { id: 10, name: 'Spitz', image: spitzImage },
    { id: 11, name: 'Boxer', image: boxerImage },
    { id: 12, name: 'Paws', image: pawsImage },
  ];
  

const DogHouseCard = () => {
    let defaultMatrix = [[11, 7, 6, 4, 1], [4, 6, 5, 4, 6], [3, 1, 3, 3, 4]]

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
            const url = `http://localhost:8081/spin-slot-v2/?spinBet=${bet}`;
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
        <Wrapper>
        <Header username={username} balance={balance} handleLogout={() => {
            localStorage.removeItem('token');
            window.location.reload();
          }} />
        <InnerWrapper>
          <Title>Dog House</Title>
          <BalanceText>Ваш баланс: {balance} UAH</BalanceText>
          <BetInput
            type="number"
            value={bet}
            onChange={e => setBet(Number(e.target.value))}
          />
          <Button onClick={handleSpin} disabled={isSpinning}>
            {isSpinning ? 'Spinning...' : 'Крутить'}
          </Button>
          <MusicButton onClick={toggleMusic}>
            {isMusicPlaying ? 'Отключить музыку' : 'Включить музыку'}
          </MusicButton>
          <GameBoard isSpinning={isSpinning}>
            {gameBoard.flat().map((symbolId, index) => {
              const symbol = symbols.find(s => s.id === symbolId);
              return (
                <Symbol key={index} color={symbol.color}>
                  <img src={symbol.image} alt={symbol.name} style={{ width: '100%', height: '100%' }} />
                </Symbol>
              );
            })}
          </GameBoard>
        </InnerWrapper>
        <audio ref={audioRef} src={backgroundMusic} />
      </Wrapper>
    );
}

export default DogHouseCard;