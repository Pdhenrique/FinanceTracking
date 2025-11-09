import React from 'react';
import styled from 'styled-components';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, ResponsiveContainer } from 'recharts';
import type { Transaction } from '../../types';

interface TransactionTrendsChartProps {
  transactions: Transaction[];
}

const TransactionTrendsChart: React.FC<TransactionTrendsChartProps> = ({ transactions }) => {
  // Processar dados para o gráfico
  const processData = () => {
    // Agrupar transações por data
    const groupedByDate = transactions.reduce((acc, transaction) => {
      const rawDate = transaction.release_date || transaction.accounting_date || '';
      if (!rawDate) {
        return acc;
      }
      const date = rawDate.substring(0, 10); // Formato YYYY-MM-DD
      if (!acc[date]) {
        acc[date] = {
          date,
          income: 0,
          expense: 0,
        };
      }

      acc[date].income += Number(transaction.income || 0);
      acc[date].expense += Number(transaction.expense || 0);

      return acc;
    }, {} as Record<string, { date: string; income: number; expense: number }>);

    // Converter para array e ordenar por data
    return Object.values(groupedByDate).sort((a, b) => {
      return new Date(a.date).getTime() - new Date(b.date).getTime();
    });
  };

  const data = processData();

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
    }).format(date);
  };

  return (
    <ChartContainer>
      <ChartHeader>
        <ChartTitle>Evolução Financeira</ChartTitle>
      </ChartHeader>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart data={data} margin={{ top: 5, right: 30, left: 20, bottom: 5 }}>
          <CartesianGrid strokeDasharray="3 3" stroke="var(--border)" />
          <XAxis
            dataKey="date"
            tickFormatter={formatDate}
            stroke="var(--text-light)"
          />
          <YAxis
            tickFormatter={(value) => `R$${value}`}
            stroke="var(--text-light)"
          />
          <Tooltip
            formatter={(value) => [formatCurrency(Number(value)), '']}
            labelFormatter={(label) => formatDate(String(label))}
          />
          <Line
            type="monotone"
            dataKey="income"
            stroke="var(--success)"
            name="Entradas"
            strokeWidth={2}
            dot={{ r: 4 }}
            activeDot={{ r: 6 }}
          />
          <Line
            type="monotone"
            dataKey="expense"
            stroke="var(--danger)"
            name="Saídas"
            strokeWidth={2}
            dot={{ r: 4 }}
            activeDot={{ r: 6 }}
          />
        </LineChart>
      </ResponsiveContainer>
    </ChartContainer>
  );
};

const ChartContainer = styled.div`
  background-color: var(--background-light);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 4px 12px var(--shadow);
  height: 100%;
`;

const ChartHeader = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
`;

const ChartTitle = styled.h3`
  font-size: 16px;
  font-weight: 500;
  color: var(--text-light);
  margin: 0;
`;

export default TransactionTrendsChart;
