package com.example.server0.model;



public class Candidate {
private Integer CandidateId;
private String CandidateName;
private String CandidateDetails;
private byte[] CandidatePhoto;

    public Integer getCandidateId() {
        return CandidateId;
    }

    public void setCandidateId(Integer candidateId) {
        CandidateId = candidateId;
    }

    public String getCandidateName() {
        return CandidateName;
    }

    public void setCandidateName(String candidateName) {
        CandidateName = candidateName;
    }

    public String getCandidateDetails() {
        return CandidateDetails;
    }

    public void setCandidateDetails(String candidateDetails) {
        CandidateDetails = candidateDetails;
    }

    public byte[] getCandidatePhoto() {
        return CandidatePhoto;
    }

    public void setCandidatePhoto(byte[] candidatePhoto) {
        CandidatePhoto = candidatePhoto;
    }
}

