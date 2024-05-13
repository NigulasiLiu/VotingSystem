package com.example.server1.model;

public class Candidate {
    private Integer candidateId;
    private String candidateName;
    private String candidateDetails;
    private byte[] candidatePhoto;

    public Integer getCandidateId() {
        return candidateId;
    }

    public void setCandidateId(Integer candidateId) {
        this.candidateId = candidateId;
    }

    public String getCandidateName() {
        return candidateName;
    }

    public void setCandidateName(String candidateName) {
        this.candidateName = candidateName;
    }

    public String getCandidateDetails() {
        return candidateDetails;
    }

    public void setCandidateDetails(String candidateDetails) {
        this.candidateDetails = candidateDetails;
    }

    public byte[] getCandidatePhoto() {
        return candidatePhoto;
    }

    public void setCandidatePhoto(byte[] candidatePhoto) {
        this.candidatePhoto = candidatePhoto;
    }
}

