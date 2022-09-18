#include <stdio.h>
#include <iostream>
#include <assert.h>
#include "circom.hpp"
#include "calcwit.hpp"
void RangeProof_0_create(uint soffset,uint coffset,Circom_CalcWit* ctx,std::string componentName,uint componentFather);
void RangeProof_0_run(uint ctx_index,Circom_CalcWit* ctx);
Circom_TemplateFunction _functionTable[1] = { 
RangeProof_0_run };
Circom_TemplateFunction _functionTableParallel[1] = { 
NULL };
uint get_main_input_signal_start() {return 2;}

uint get_main_input_signal_no() {return 3;}

uint get_total_signal_no() {return 7;}

uint get_number_of_components() {return 1;}

uint get_size_of_input_hashmap() {return 256;}

uint get_size_of_witness() {return 6;}

uint get_size_of_constants() {return 4;}

uint get_size_of_io_map() {return 0;}

// function declarations
// template declarations
void RangeProof_0_create(uint soffset,uint coffset,Circom_CalcWit* ctx,std::string componentName,uint componentFather){
ctx->componentMemory[coffset].templateId = 0;
ctx->componentMemory[coffset].templateName = "RangeProof";
ctx->componentMemory[coffset].signalStart = soffset;
ctx->componentMemory[coffset].inputCounter = 3;
ctx->componentMemory[coffset].componentName = componentName;
ctx->componentMemory[coffset].idFather = componentFather;
ctx->componentMemory[coffset].subcomponents = new uint[0];
}

void RangeProof_0_run(uint ctx_index,Circom_CalcWit* ctx){
FrElement* signalValues = ctx->signalValues;
u64 mySignalStart = ctx->componentMemory[ctx_index].signalStart;
std::string myTemplateName = ctx->componentMemory[ctx_index].templateName;
std::string myComponentName = ctx->componentMemory[ctx_index].componentName;
u64 myFather = ctx->componentMemory[ctx_index].idFather;
u64 myId = ctx_index;
u32* mySubcomponents = ctx->componentMemory[ctx_index].subcomponents;
bool* mySubcomponentsParallel = ctx->componentMemory[ctx_index].subcomponentsParallel;
FrElement* circuitConstants = ctx->circuitConstants;
std::string* listOfTemplateMessages = ctx->listOfTemplateMessages;
FrElement expaux[5];
FrElement lvar[0];
uint sub_component_aux;
{
PFrElement aux_dest = &signalValues[mySignalStart + 4];
// load src
Fr_mul(&expaux[0],&signalValues[mySignalStart + 2],&signalValues[mySignalStart + 3]); // line circom 12
// end load src
Fr_copy(aux_dest,&expaux[0]);
}
Fr_mul(&expaux[2],&signalValues[mySignalStart + 2],&signalValues[mySignalStart + 1]); // line circom 13
Fr_leq(&expaux[0],&signalValues[mySignalStart + 4],&expaux[2]); // line circom 13
if(Fr_isTrue(&expaux[0])){
{
PFrElement aux_dest = &signalValues[mySignalStart + 5];
// load src
// end load src
Fr_copy(aux_dest,&circuitConstants[0]);
}
}else{
{
PFrElement aux_dest = &signalValues[mySignalStart + 5];
// load src
// end load src
Fr_copy(aux_dest,&circuitConstants[1]);
}
}
{
PFrElement aux_dest = &signalValues[mySignalStart + 0];
// load src
Fr_mul(&expaux[1],&signalValues[mySignalStart + 5],&circuitConstants[2]); // line circom 16
Fr_add(&expaux[0],&expaux[1],&circuitConstants[3]); // line circom 16
// end load src
Fr_copy(aux_dest,&expaux[0]);
}
}

void run(Circom_CalcWit* ctx){
RangeProof_0_create(1,0,ctx,"main",0);
RangeProof_0_run(0,ctx);
}

