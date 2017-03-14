#############################################################
# description: determine if a given configuration is valid  # 
# game_mat: a matrix representing the combined board state  #
# black_mat: a matrix representing only black pieces        #
# white_mat: a matrix representing white's pieces           #
# turn     : 0 if black's turn 1 if white's turn            #
# returns: 1 if the configuration is valid, 0 if not        #
#############################################################

import numpy.linalg as la
from numpy.linalg import svd
import numpy as np
from rank_nullspace import rank,nullspace

def valid_move(game_mat, black_mat, white_mat, turn):
    BOARD_SIZE = game_mat.shape[0]
    WHITE_MAT_WIDTH = white_mat.shape[0]
    BLACK_MAT_WIDTH = black_mat.shape[0]

    null_black= nullspace(black_mat)
    null_white = nullspace(white_mat)
    black_dict = {}
    white_dict = {}
    white_captured = 0
    black_captured = 0

    #This is just some linear algebra magic to efficently look for captured components
    for i in range(len(null_black[0])):
        col = null_black[:,i]
        black_dict[i] = np.nonzero(col)
    for i in range(len(null_white[0])):
        col = null_white[:,i]
        white_dict[i] = np.nonzero(col)

    diag = np.diag(game_mat)
    for i in white_dict:
        arr = white_dict[i][0]
        for j in range(len(arr)):
            val = arr[j]
            if diag[val] != 4:
                break;
            if j == len(arr)- 1:
                white_captured += 1

    for i in black_dict:
        arr = black_dict[i][0]
        for j in range(len(arr)):
            val = arr[j]
            if diag[val + WHITE_MAT_WIDTH] != 4:
                break;
            if j == len(arr)- 1:
                black_captured += 1

    #this is the actual valid move check
    if white_captured + black_captured == 1:
        if white_captured == 1 and turn == 1:
            return 0
        if black_captured == 1 and turn == 0:
            return 0
    return 1
