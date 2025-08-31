import React from 'react';
import styled from 'styled-components';
import { FiArrowUpRight, FiArrowDownLeft, FiRepeat } from 'react-icons/fi';
import type { Transaction } from '../../types';

interface TransactionListProps {
  transactions: Transaction[];
  onRemoveTransaction: (id: string) => void;
}

const TransactionList: React.FC<TransactionListProps> = ({ transactions, onRemoveTransaction }) => {
  // Formatar valores monetários
  const formatCurrency = (value: number) => {
    return new Intl.NumberFormat('pt-BR', {
      style: 'currency',
      currency: 'BRL',
    }).format(value);
  };

  // Formatar datas
  const formatDate = (dateStr: string) => {
    const date = new Date(dateStr);
    return new Intl.DateTimeFormat('pt-BR', {
      day: '2-digit',
      month: '2-digit',
      year: 'numeric',
    }).format(date);
  };

  // Obter ícone com base no tipo de transação
  const getTransactionIcon = (type: string) => {
    switch (type) {
      case 'income':
        return <FiArrowUpRight size={18} />;
      case 'expense':
        return <FiArrowDownLeft size={18} />;
      case 'transfer':
        return <FiRepeat size={18} />;
      default:
        return null;
    }
  };

  return (
    <ListContainer>
      <ListHeader>
        <ListTitle>Transações Recentes</ListTitle>
      </ListHeader>

      {transactions.length === 0 ? (
        <EmptyState>Nenhuma transação encontrada.</EmptyState>
      ) : (
        <TransactionsTable>
          <thead>
            <tr>
              <TableHeader>Descrição</TableHeader>
              <TableHeader>Categoria</TableHeader>
              <TableHeader>Data</TableHeader>
              <TableHeader align="right">Valor</TableHeader>
              <TableHeader></TableHeader>
            </tr>
          </thead>
          <tbody>
            {transactions.map((transaction) => (
              <TableRow key={transaction.id}>
                <TableCell>
                  <TransactionInfo>
                    <IconContainer type={transaction.type}>
                      {getTransactionIcon(transaction.type)}
                    </IconContainer>
                    <span>{transaction.description}</span>
                  </TransactionInfo>
                </TableCell>
                <TableCell>
                  <CategoryBadge>{transaction.category}</CategoryBadge>
                </TableCell>
                <TableCell>{formatDate(transaction.date)}</TableCell>
                <TableCell align="right">
                  <TransactionAmount type={transaction.type}>
                    {transaction.type === 'expense' ? '- ' : ''}
                    {formatCurrency(transaction.amount)}
                  </TransactionAmount>
                </TableCell>
                <TableCell align="right">
                  <DeleteButton onClick={() => onRemoveTransaction(transaction.id)}>
                    Excluir
                  </DeleteButton>
                </TableCell>
              </TableRow>
            ))}
          </tbody>
        </TransactionsTable>
      )}
    </ListContainer>
  );
};

const ListContainer = styled.div`
  background-color: var(--background-light);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 12px var(--shadow);
  height: 100%;
  overflow: hidden;
`;

const ListHeader = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
`;

const ListTitle = styled.h3`
  font-size: 18px;
  font-weight: 600;
  color: var(--text);
  margin: 0;
`;

const EmptyState = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  height: 200px;
  color: var(--text-light);
  font-size: 16px;
`;

const TransactionsTable = styled.table`
  width: 100%;
  border-collapse: collapse;
`;

const TableHeader = styled.th<{ align?: string }>`
  padding: 12px 16px;
  text-align: ${(props) => props.align || 'left'};
  font-weight: 500;
  font-size: 14px;
  color: var(--text-light);
  border-bottom: 1px solid var(--border);
`;

const TableRow = styled.tr`
  &:hover {
    background-color: var(--border-light);
  }
`;

const TableCell = styled.td<{ align?: string }>`
  padding: 16px;
  text-align: ${(props) => props.align || 'left'};
  border-bottom: 1px solid var(--border-light);
  font-size: 14px;
`;

const TransactionInfo = styled.div`
  display: flex;
  align-items: center;
`;

const IconContainer = styled.div<{ type: string }>`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  border-radius: 50%;
  margin-right: 12px;
  
  ${({ type }) => {
    switch (type) {
      case 'income':
        return `
          background-color: rgba(16, 185, 129, 0.1);
          color: var(--success);
        `;
      case 'expense':
        return `
          background-color: rgba(239, 68, 68, 0.1);
          color: var(--danger);
        `;
      case 'transfer':
        return `
          background-color: rgba(245, 158, 11, 0.1);
          color: var(--warning);
        `;
      default:
        return '';
    }
  }}
`;

const CategoryBadge = styled.span`
  display: inline-block;
  padding: 4px 8px;
  background-color: var(--border-light);
  border-radius: 4px;
  font-size: 12px;
  font-weight: 500;
  color: var(--text-light);
`;

const TransactionAmount = styled.span<{ type: string }>`
  font-weight: 600;
  
  ${({ type }) => {
    switch (type) {
      case 'income':
        return `color: var(--success);`;
      case 'expense':
        return `color: var(--danger);`;
      case 'transfer':
        return `color: var(--warning);`;
      default:
        return `color: var(--text);`;
    }
  }}
`;

const DeleteButton = styled.button`
  background-color: transparent;
  border: none;
  color: var(--text-light);
  font-size: 12px;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
  transition: all 0.2s ease;
  
  &:hover {
    background-color: rgba(239, 68, 68, 0.1);
    color: var(--danger);
  }
`;

export default TransactionList;