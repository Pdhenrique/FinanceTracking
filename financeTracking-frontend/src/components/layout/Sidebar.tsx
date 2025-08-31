import React from 'react';
import styled from 'styled-components';
import { FiHome, FiPieChart, FiDollarSign, FiCreditCard, FiSettings, FiLogOut } from 'react-icons/fi';

interface SidebarProps {
  activeItem: string;
  onItemClick: (item: string) => void;
}

const Sidebar: React.FC<SidebarProps> = ({ activeItem, onItemClick }) => {
  const menuItems = [
    { id: 'dashboard', icon: <FiHome size={20} />, label: 'Dashboard' },
    { id: 'transactions', icon: <FiDollarSign size={20} />, label: 'Transações' },
    { id: 'analytics', icon: <FiPieChart size={20} />, label: 'Análises' },
    { id: 'cards', icon: <FiCreditCard size={20} />, label: 'Cartões' },
    { id: 'settings', icon: <FiSettings size={20} />, label: 'Configurações' },
  ];

  return (
    <SidebarContainer>
      <Logo>
        <LogoText>FinTech</LogoText>
      </Logo>
      <MenuItems>
        {menuItems.map((item) => (
          <MenuItem
            key={item.id}
            active={activeItem === item.id ? 'true' : 'false'}
            onClick={() => onItemClick(item.id)}
          >
            <MenuIcon>{item.icon}</MenuIcon>
            <MenuLabel>{item.label}</MenuLabel>
          </MenuItem>
        ))}
      </MenuItems>
      <LogoutButton>
        <MenuIcon>
          <FiLogOut size={20} />
        </MenuIcon>
        <MenuLabel>Sair</MenuLabel>
      </LogoutButton>
    </SidebarContainer>
  );
};

const SidebarContainer = styled.div`
  display: flex;
  flex-direction: column;
  width: 250px;
  height: 100vh;
  background-color: var(--background-light);
  border-right: 1px solid var(--border);
  padding: 20px 0;
`;

const Logo = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 20px 0;
  margin-bottom: 20px;
`;

const LogoText = styled.h1`
  font-size: 24px;
  font-weight: 700;
  color: var(--primary);
`;

const MenuItems = styled.div`
  display: flex;
  flex-direction: column;
  flex: 1;
`;

const MenuItem = styled.div<{ active: string }>`
  display: flex;
  align-items: center;
  padding: 15px 25px;
  cursor: pointer;
  transition: all 0.2s ease;
  background-color: ${(props) => (props.active === 'true' ? 'var(--primary-light)' : 'transparent')};
  color: ${(props) => (props.active === 'true' ? 'white' : 'var(--text)')};
  border-left: ${(props) => (props.active === 'true' ? '4px solid var(--primary)' : '4px solid transparent')};

  &:hover {
    background-color: ${(props) => (props.active === 'true' ? 'var(--primary-light)' : 'var(--border-light)')};
  }
`;

const MenuIcon = styled.div`
  margin-right: 15px;
  display: flex;
  align-items: center;
`;

const MenuLabel = styled.span`
  font-size: 16px;
  font-weight: 500;
`;

const LogoutButton = styled.div`
  display: flex;
  align-items: center;
  padding: 15px 25px;
  cursor: pointer;
  transition: all 0.2s ease;
  color: var(--text-light);
  margin-top: auto;
  border-top: 1px solid var(--border);

  &:hover {
    color: var(--danger);
  }
`;

export default Sidebar;