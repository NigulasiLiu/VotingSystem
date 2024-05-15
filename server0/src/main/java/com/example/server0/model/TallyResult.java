package com.example.server0.model;

import java.math.BigInteger;
import java.util.List;
import java.util.Map;

public class TallyResult {
    private Map<Integer, List<String>> vwTally;
    private Map<Integer,List<BigInteger[]>> outputsTally;
    public void setVwTally(Map<Integer, List<String>> vwTally) {
        this.vwTally = vwTally;
    }
    public void setOutputsTally(Map<Integer,List<BigInteger[]>> outputsTally) {
        this.outputsTally = outputsTally;
    }
    
}
