package com.personal.financeTracking.enums;

import com.fasterxml.jackson.annotation.JsonCreator;
import com.fasterxml.jackson.annotation.JsonValue;
import com.personal.financeTracking.transaction.entities.Transaction;

public enum TransactionType {
    INCOME,
    EXPENSE,
    TRANSFER;

    @JsonCreator
    public static TransactionType fromString(String value){
      try {
        return TransactionType.valueOf(value.toUpperCase());
      } catch (IllegalArgumentException e) {
        throw new RuntimeException("Invalid transaction type:" + value);
      }
  }

  @JsonValue
  public String toValue(){ return this.name();}
}
