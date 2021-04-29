#ifndef KCD_TOKENIZER_H
#define KCD_TOKENIZER_H

#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

typedef struct {
	uint64_t location;
	char token;
} token_t;

token_t* tokenizeFile(char* file, char** contentsPTR, uint64_t* numTokens, uint64_t* flen);

#endif