import { useState, useEffect } from "react";
import type { Transaction, BalanceSummary, CategorySummary } from "../types";

// Dados de exemplo para demonstração
const mockTransactions: Transaction[] = [
  {
    id: "1",
    description: "Salário",
    amount: 5000,
    date: "2023-08-01",
    category: "Renda",
    type: "income",
  },
  {
    id: "2",
    description: "Aluguel",
    amount: 1200,
    date: "2023-08-05",
    category: "Moradia",
    type: "expense",
  },
  {
    id: "3",
    description: "Supermercado",
    amount: 450,
    date: "2023-08-10",
    category: "Alimentação",
    type: "expense",
  },
  {
    id: "4",
    description: "Freelance",
    amount: 1200,
    date: "2023-08-15",
    category: "Renda",
    type: "income",
  },
  {
    id: "5",
    description: "Transferência para poupança",
    amount: 800,
    date: "2023-08-20",
    category: "Transferência",
    type: "transfer",
  },
  {
    id: "6",
    description: "Conta de luz",
    amount: 150,
    date: "2023-08-22",
    category: "Utilidades",
    type: "expense",
  },
  {
    id: "7",
    description: "Conta de internet",
    amount: 120,
    date: "2023-08-22",
    category: "Utilidades",
    type: "expense",
  },
  {
    id: "8",
    description: "Jantar fora",
    amount: 180,
    date: "2023-08-25",
    category: "Lazer",
    type: "expense",
  },
];

export const useTransactions = () => {
  const [transactions, setTransactions] =
    useState<Transaction[]>(mockTransactions);
  const [balanceSummary, setBalanceSummary] = useState<BalanceSummary>({
    totalBalance: 0,
    income: 0,
    expenses: 0,
  });
  const [categorySummary, setCategorySummary] = useState<CategorySummary[]>([]);

  // Calcular resumo de saldo
  useEffect(() => {
    const income = transactions
      .filter((transaction) => transaction.type === "income")
      .reduce((acc, transaction) => acc + transaction.amount, 0);

    const expenses = transactions
      .filter((transaction) => transaction.type === "expense")
      .reduce((acc, transaction) => acc + transaction.amount, 0);

    const totalBalance = income - expenses;

    setBalanceSummary({ totalBalance, income, expenses });
  }, [transactions]);

  // Calcular resumo por categoria
  useEffect(() => {
    const expensesByCategory = transactions
      .filter((transaction) => transaction.type === "expense")
      .reduce((acc, transaction) => {
        const { category, amount } = transaction;
        if (!acc[category]) {
          acc[category] = 0;
        }
        acc[category] += amount;
        return acc;
      }, {} as Record<string, number>);

    const totalExpenses = Object.values(expensesByCategory).reduce(
      (acc, amount) => acc + amount,
      0
    );

    const summary = Object.entries(expensesByCategory).map(
      ([category, amount]) => ({
        category,
        amount,
        percentage: totalExpenses > 0 ? (amount / totalExpenses) * 100 : 0,
      })
    );

    setCategorySummary(summary);
  }, [transactions]);

  // Adicionar nova transação
  const addTransaction = (transaction: Omit<Transaction, "id">) => {
    const newTransaction = {
      ...transaction,
      id: Math.random().toString(36).substring(2, 9),
    };

    setTransactions((prev) => [newTransaction, ...prev]);
  };

  // Remover transação
  const removeTransaction = (id: string) => {
    setTransactions((prev) =>
      prev.filter((transaction) => transaction.id !== id)
    );
  };

  return {
    transactions,
    balanceSummary,
    categorySummary,
    addTransaction,
    removeTransaction,
  };
};
