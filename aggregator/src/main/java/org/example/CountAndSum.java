package org.example;


import com.fasterxml.jackson.annotation.JsonProperty;

public class CountAndSum {
    @JsonProperty("sum")
    public long sum;

    @JsonProperty("count")
    public long count;

    public CountAndSum() {
    }

    public CountAndSum(long sum, long count) {
        this.sum = sum;
        this.count = count;
    }
}
