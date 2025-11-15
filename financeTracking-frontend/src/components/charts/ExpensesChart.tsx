import React from 'react';
import styled from 'styled-components';
import { PieChart, Pie, Cell, ResponsiveContainer, Tooltip, Legend } from 'recharts';
import type { CategorySummary } from '../../types';

interface ExpensesChartProps {
  data: CategorySummary[];
}

const COLORS = ['#0088FE', '#00C49F', '#FFBB28', '#FF8042', '#8884D8', '#82CA9D', '#FF6B6B', '#6B66FF'];

const ExpensesChart: React.FC<ExpensesChartProps> = ({ data }) => {
  return (
    <ChartContainer>
      <ChartHeader>
        <ChartTitle>Despesas por Categoria</ChartTitle>
      </ChartHeader>
      <ResponsiveContainer width="100%" height={300}>
        <PieChart>
          <Pie
            data={data}
            cx="50%"
            cy="50%"
            labelLine={false}
            outerRadius={100}
            fill="#8884d8"
            dataKey="amount"
            nameKey="category"
            label={({ name, percent }) => `${name}: ${(percent * 100).toFixed(0)}%`}
          >
            {data.map((_, index) => (
              <Cell key={`cell-${index}`} fill={COLORS[index % COLORS.length]} />
            ))}
          </Pie>
          <Tooltip
            formatter={(value) => [
              new Intl.NumberFormat('pt-BR', {
                style: 'currency',
                currency: 'BRL',
              }).format(Number(value)),
              'Valor',
            ]}
          />
          <Legend />
        </PieChart>
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

export default ExpensesChart;