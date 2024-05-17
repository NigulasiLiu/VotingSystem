package com.example.server0.model;

public class Result {
    private int id;
    private Integer voteId;
    private byte[] outputs0;
    private String vw0;
    private byte[] outputs1;
    private String vw1;

    public byte[] getOutputs1() {
        return outputs1;
    }

    public void setOutputs1(byte[] outputs1) {
        this.outputs1 = outputs1;
    }

    public String getVw1() {
        return vw1;
    }

    public void setVw1(String vw1) {
        this.vw1 = vw1;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public Integer getVoteId() {
        return voteId;
    }

    public void setVoteId(Integer voteId) {
        this.voteId = voteId;
    }

    public byte[] getOutputs0() {
        return outputs0;
    }

    public void setOutputs0(byte[] outputs0) {
        this.outputs0 = outputs0;
    }

    public String getVw0() {
        return vw0;
    }

    public void setVw0(String vw0) {
        this.vw0 = vw0;
    }
}
