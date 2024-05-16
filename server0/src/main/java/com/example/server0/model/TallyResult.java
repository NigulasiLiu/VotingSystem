package com.example.server0.model;

import java.math.BigInteger;
import java.util.List;
import java.util.Map;

public class TallyResult {
    private String[] vwTally;
    private Map<Integer,List<BigInteger[]>> outputsTally;

    public String[] getVwTally() {
        return vwTally;
    }

    public Map<Integer, List<BigInteger[]>> getOutputsTally() {
        return outputsTally;
    }

    public void setVwTally(String[] vwTally) {
        this.vwTally = vwTally;
    }
    public void setOutputsTally(Map<Integer,List<BigInteger[]>> outputsTally) {
        this.outputsTally = outputsTally;
    }
    
}
