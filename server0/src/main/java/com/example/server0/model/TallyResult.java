package com.example.server0.model;

import java.util.List;
import java.util.Map;

public class TallyResult {
    private Map<Integer, List<byte[]>> vwTally;
    private Map<Integer,List<double[]>> outputsTally;
    public void setVwTally(Map<Integer, List<byte[]>> vwTally) {
        this.vwTally = vwTally;
    }
    public void setOutputsTally(Map<Integer, List<double[]>> outputsTally) {
        this.outputsTally = outputsTally;
    }
    
}
