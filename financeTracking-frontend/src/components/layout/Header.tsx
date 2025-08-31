import React from 'react';
import styled from 'styled-components';
import { FiBell, FiSearch, FiUser, FiMoon, FiSun } from 'react-icons/fi';
import { useTheme } from '../../contexts/ThemeContext';

interface HeaderProps {
  userName: string;
}

const Header: React.FC<HeaderProps> = ({ userName }) => {
  const { theme, toggleTheme } = useTheme();
  
  return (
    <HeaderContainer>
      <SearchContainer>
        <SearchIcon>
          <FiSearch size={18} />
        </SearchIcon>
        <SearchInput placeholder="Buscar..." />
      </SearchContainer>
      
      <RightSection>
        <ThemeToggleButton onClick={toggleTheme}>
          {theme === 'light' ? <FiMoon size={20} /> : <FiSun size={20} />}
        </ThemeToggleButton>
        
        <NotificationButton>
          <FiBell size={20} />
          <NotificationBadge>3</NotificationBadge>
        </NotificationButton>
        
        <UserProfile>
          <UserAvatar>
            <FiUser size={20} />
          </UserAvatar>
          <UserName>{userName}</UserName>
        </UserProfile>
      </RightSection>
    </HeaderContainer>
  );
};

const HeaderContainer = styled.header`
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 15px 30px;
  background-color: var(--background-light);
  border-bottom: 1px solid var(--border);
  height: 70px;
`;

const SearchContainer = styled.div`
  display: flex;
  align-items: center;
  background-color: var(--background);
  border-radius: 8px;
  padding: 8px 15px;
  width: 300px;
`;

const SearchIcon = styled.div`
  color: var(--text-light);
  margin-right: 10px;
`;

const SearchInput = styled.input`
  border: none;
  background: transparent;
  outline: none;
  color: var(--text);
  font-size: 14px;
  width: 100%;
  
  &::placeholder {
    color: var(--text-light);
  }
`;

const RightSection = styled.div`
  display: flex;
  align-items: center;
`;

const NotificationButton = styled.button`
  position: relative;
  background: transparent;
  border: none;
  color: var(--text);
  margin-right: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  padding: 8px;
  border-radius: 50%;
  transition: all 0.2s ease;
  
  &:hover {
    background-color: var(--border-light);
  }
`;

const NotificationBadge = styled.span`
  position: absolute;
  top: 0;
  right: 0;
  background-color: var(--danger);
  color: white;
  font-size: 10px;
  font-weight: bold;
  width: 16px;
  height: 16px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const UserProfile = styled.div`
  display: flex;
  align-items: center;
  cursor: pointer;
`;

const UserAvatar = styled.div`
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background-color: var(--primary-light);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  margin-right: 10px;
`;

const UserName = styled.span`
  font-weight: 500;
  font-size: 14px;
  color: var(--text);
`;

const ThemeToggleButton = styled.button`
  background: transparent;
  border: none;
  color: var(--text);
  margin-right: 20px;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  padding: 8px;
  border-radius: 50%;
  transition: all 0.2s ease;
  
  &:hover {
    background-color: var(--border-light);
  }
`;

export default Header;