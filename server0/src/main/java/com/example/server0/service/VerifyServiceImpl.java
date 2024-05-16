package com.example.server0.service;


import com.example.server0.controller.TallyingController;
import com.example.server0.model.TallyResult;
import com.example.server0.model.VoteData;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.math.BigInteger;
import java.util.List;

@Service
public class VerifyServiceImpl implements VerifyService {
    @Autowired
    private TallyingController tallyingController;
    TallyResult result=new TallyResult();
    String[] vwResult= result.getVwTally();
    public String[] computeTally(String[] vwResult0,String[] vwResult1,int lamda) {
        List<VoteData> votesList=tallyingController.getList();
        int nums=votesList.size();
        String[] nHonest=new String[nums];
        for (int i = 0; i < nHonest.length; i++) {
            String pi0=vwResult0[i].substring(0,lamda+lamda);
            String Val0=vwResult0[i].substring(lamda+lamda,3*lamda+2);
            String pi1=vwResult1[i].substring(0,lamda+lamda);
            String Val1=vwResult1[i].substring(lamda+lamda,3*lamda+2);
            BigInteger val0=new BigInteger(Val0);
            BigInteger val1=new BigInteger(Val1);
            if(pi0==pi1 && val0.add(val1)==1){
                nHonest[i]="0";
            }else{
                nHonest[i]="1";
            }

        }
    }
}