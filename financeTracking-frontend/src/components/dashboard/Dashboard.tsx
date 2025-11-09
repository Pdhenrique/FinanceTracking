import React from 'react';
import styled from 'styled-components';
import BalanceCard from './BalanceCard';
import TransactionList from '../transactions/TransactionList';
import AddTransactionForm from '../transactions/AddTransactionForm';
import ExpensesChart from '../charts/ExpensesChart';
import TransactionTrendsChart from '../charts/TransactionTrendsChart';
import { useTransactions } from '../../hooks/useTransactions';

const Dashboard: React.FC = () => {
  const { 
    transactions, 
    balanceSummary, 
    categorySummary,
    addTransaction, 
    removeTransaction 
  } = useTransactions();

  return (
    <DashboardContainer>
      <TopSection>
        <BalanceCardWrapper>
          <BalanceCard 
            totalBalance={balanceSummary.totalBalance}
            income={balanceSummary.income}
            expenses={balanceSummary.expenses}
          />
        </BalanceCardWrapper>
        
        <FormWrapper>
          <AddTransactionForm onAddTransaction={addTransaction} />
        </FormWrapper>
      </TopSection>
      
      <ChartsSection>
        <ChartWrapper>
          <ExpensesChart data={categorySummary} />
        </ChartWrapper>
        
        <ChartWrapper>
          <TransactionTrendsChart transactions={transactions} />
        </ChartWrapper>
      </ChartsSection>
      
      <TransactionsSection>
        <TransactionList 
          transactions={transactions}
          onRemoveTransaction={removeTransaction}
        />
      </TransactionsSection>
    </DashboardContainer>
  );
};

const DashboardContainer = styled.div`
  display: flex;
  flex-direction: column;
  gap: 24px;
`;

const TopSection = styled.div`
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 24px;
  
  @media (max-width: 768px) {
    grid-template-columns: 1fr;
  }
`;

const BalanceCardWrapper = styled.div`
  flex: 2;
`;

const FormWrapper = styled.div`
  flex: 1;
`;

const ChartsSection = styled.div`
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 24px;
  
  @media (max-width: 768px) {
    grid-template-columns: 1fr;
  }
`;

const ChartWrapper = styled.div`
  height: 350px;
`;

const TransactionsSection = styled.div`
  margin-top: 24px;
`;

export default Dashboard;