package com.personal.financeTracking.enums;

public enum AccountType {
    CONTA_CORRENTE,
    POUPANCA,
    CARTAO_CREDITO,
    INVESTIMENTO;

    public static AccountType fromString(String value) {
        try {
            return AccountType.valueOf(value.toUpperCase());
        } catch (IllegalArgumentException e) {
            throw new RuntimeException("Invalid account type: " + value);
        }
    }
}
