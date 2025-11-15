import React, { useState } from 'react';
import styled from 'styled-components';
import Sidebar from './Sidebar';
import Header from './Header';

interface LayoutProps {
  children: React.ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  const [activeItem, setActiveItem] = useState('dashboard');

  return (
    <LayoutContainer>
      <Sidebar activeItem={activeItem} onItemClick={setActiveItem} />
      <MainContent>
        <Header userName="Pedro Silva" />
        <ContentArea>{children}</ContentArea>
      </MainContent>
    </LayoutContainer>
  );
};

const LayoutContainer = styled.div`
  display: flex;
  height: 100vh;
  overflow: hidden;
`;

const MainContent = styled.div`
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
`;

const ContentArea = styled.main`
  flex: 1;
  padding: 25px;
  overflow-y: auto;
  background-color: var(--background);
`;

export default Layout;