package com.example.server0.model;

public class Voter {
    private Integer VoterId;
    private String VoterName;
    private String password;
    private String VoterType;

    public Integer getVoterId() {
        return VoterId;
    }

    public void setVoterId(Integer VoterId) {
        this.VoterId = VoterId;
    }

    public String getVoterName() {
        return VoterName;
    }

    public void setVoterName(String VoterName) {
        this.VoterName = VoterName;
    }

    public String getVoterType() {
        return VoterType;
    }

    public void setVoterTyoe(String VoterType) {
        this.VoterType = VoterType;
    }

    public String getPassword() {
        return password;
    }

    public void setPassword(String password) {
        this.password = password;
    }
}