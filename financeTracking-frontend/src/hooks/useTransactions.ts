import { useState, useEffect } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import axios from "axios";
import type { Transaction, BalanceSummary, CategorySummary } from "../types";

const API_URL = "http://localhost:8080/v1/transactions";

// Funções de API
const fetchTransactions = async (): Promise<Transaction[]> => {
  const response = await axios.get(API_URL);
  return response.data;
};

const addTransactionAPI = async (
  transaction: Omit<Transaction, "id">
): Promise<Transaction> => {
  const response = await axios.post(API_URL, transaction);
  return response.data;
};

const removeTransactionAPI = async (id: string): Promise<void> => {
  await axios.delete(`${API_URL}/${id}`);
  return;
};

export const useTransactions = () => {
  const queryClient = useQueryClient();
  const [balanceSummary, setBalanceSummary] = useState<BalanceSummary>({
    totalBalance: 0,
    income: 0,
    expenses: 0,
  });
  const [categorySummary, setCategorySummary] = useState<CategorySummary[]>([]);

  // Buscar transações usando React Query
  const {
    data: transactions = [],
    isLoading,
    isError,
  } = useQuery({
    queryKey: ["transactions"],
    queryFn: fetchTransactions,
    staleTime: 60 * 1000, // 1 minuto
    refetchOnWindowFocus: false,
  });

  // Calcular resumo de saldo quando as transações mudarem
  useEffect(() => {
    if (!isLoading && transactions && transactions.length > 0) {
      const totalIncome = transactions.reduce(
        (acc, transaction) => acc + transaction.income,
        0
      );
      const totalExpenses = transactions.reduce(
        (acc, transaction) => acc + transaction.expense,
        0
      );
      const totalBalance = totalIncome - totalExpenses;

      setBalanceSummary({
        totalBalance,
        income: totalIncome,
        expenses: totalExpenses,
      });
    }
  }, [transactions, isLoading]);

  // Calcular resumo por categoria quando as transações mudarem
  useEffect(() => {
    if (!isLoading && transactions && transactions.length > 0) {
      // Agrupar despesas por título (usado como categoria)
      const expensesByTitle = transactions
        .filter((transaction) => transaction.expense > 0)
        .reduce((acc, transaction) => {
          const { title, expense } = transaction;
          if (!acc[title]) {
            acc[title] = 0;
          }
          acc[title] += expense;
          return acc;
        }, {} as Record<string, number>);

      const totalExpenses = Object.values(expensesByTitle).reduce(
        (acc, amount) => acc + amount,
        0
      );

      const summary = Object.entries(expensesByTitle).map(
        ([title, amount]) => ({
          category: title, // Mantendo a interface CategorySummary
          amount,
          percentage: totalExpenses > 0 ? (amount / totalExpenses) * 100 : 0,
        })
      );

      setCategorySummary(summary);
    }
  }, [transactions, isLoading]);

  // Mutation para adicionar transação
  const addTransactionMutation = useMutation({
    mutationFn: addTransactionAPI,
    onSuccess: () => {
      // Invalidar e buscar novamente a query de transações
      queryClient.invalidateQueries({ queryKey: ["transactions"] });
    },
    onError: (error) => {
      console.error("Erro ao adicionar transação:", error);
    },
  });

  // Mutation para remover transação
  const removeTransactionMutation = useMutation({
    mutationFn: removeTransactionAPI,
    onSuccess: () => {
      // Invalidar e buscar novamente a query de transações
      queryClient.invalidateQueries({ queryKey: ["transactions"] });
    },
    onError: (error) => {
      console.error("Erro ao remover transação:", error);
    },
  });

  // Adicionar nova transação
  const addTransaction = (transaction: Omit<Transaction, "id">) => {
    addTransactionMutation.mutate(transaction);
  };

  // Remover transação
  const removeTransaction = (id: string) => {
    removeTransactionMutation.mutate(id);
  };

  return {
    // Dados
    transactions,
    balanceSummary,
    categorySummary,

    // Funções
    addTransaction,
    removeTransaction,

    // Estados da query
    isLoading,
    isError,

    // Estados das mutations
    isAddingTransaction: addTransactionMutation.isPending,
    isRemovingTransaction: removeTransactionMutation.isPending,
    addTransactionError: addTransactionMutation.error,
    removeTransactionError: removeTransactionMutation.error,
  };
};
