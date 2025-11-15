import React from 'react';
import styled from 'styled-components';
import { FiArrowUp, FiArrowDown } from 'react-icons/fi';

interface BalanceCardProps {
  totalBalance: number;
  income: number;
  expenses: number;
}

const BalanceCard: React.FC<BalanceCardProps> = ({ totalBalance, income, expenses }) => {
  // Formatar valores monetários
  const formatCurrency = (value: number) => {
    return new Intl.NumberFormat('pt-BR', {
      style: 'currency',
      currency: 'BRL',
    }).format(value);
  };

  return (
    <CardContainer>
      <CardHeader>
        <CardTitle>Saldo Total</CardTitle>
      </CardHeader>
      <TotalBalance>{formatCurrency(totalBalance)}</TotalBalance>
      
      <FlowContainer>
        <FlowItem>
          <FlowIcon income="true">
            <FiArrowUp size={16} />
          </FlowIcon>
          <FlowDetails>
            <FlowLabel>Entradas</FlowLabel>
            <FlowValue income="true">{formatCurrency(income)}</FlowValue>
          </FlowDetails>
        </FlowItem>
        
        <FlowItem>
          <FlowIcon income="false">
            <FiArrowDown size={16} />
          </FlowIcon>
          <FlowDetails>
            <FlowLabel>Saídas</FlowLabel>
            <FlowValue income="false">{formatCurrency(expenses)}</FlowValue>
          </FlowDetails>
        </FlowItem>
      </FlowContainer>
    </CardContainer>
  );
};

const CardContainer = styled.div`
  background-color: var(--background-light);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 12px var(--shadow);
  transition: transform 0.2s ease;
  
  &:hover {
    transform: translateY(-5px);
  }
`;

const CardHeader = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
`;

const CardTitle = styled.h3`
  font-size: 16px;
  font-weight: 500;
  color: var(--text-light);
  margin: 0;
`;

const TotalBalance = styled.h2`
  font-size: 32px;
  font-weight: 700;
  color: var(--text);
  margin: 0 0 24px 0;
`;

const FlowContainer = styled.div`
  display: flex;
  justify-content: space-between;
  gap: 16px;
`;

const FlowItem = styled.div`
  display: flex;
  align-items: center;
  flex: 1;
`;

const FlowIcon = styled.div<{ income: string }>`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  background-color: ${(props) => (props.income === 'true' ? 'rgba(16, 185, 129, 0.1)' : 'rgba(239, 68, 68, 0.1)')};
  color: ${(props) => (props.income === 'true' ? 'var(--success)' : 'var(--danger)')};
  margin-right: 12px;
`;

const FlowDetails = styled.div`
  display: flex;
  flex-direction: column;
`;

const FlowLabel = styled.span`
  font-size: 12px;
  color: var(--text-light);
  margin-bottom: 4px;
`;

const FlowValue = styled.span<{ income: string }>`
  font-size: 16px;
  font-weight: 600;
  color: ${(props) => (props.income === 'true' ? 'var(--success)' : 'var(--danger)')};
`;

export default BalanceCard;