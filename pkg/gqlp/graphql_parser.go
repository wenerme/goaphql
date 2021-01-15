// Code generated from GraphQL.g4 by ANTLR 4.9.1. DO NOT EDIT.

package gqlp // GraphQL
import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 51, 697,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54, 4, 55,
	9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4, 60, 9,
	60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65, 9, 65,
	4, 66, 9, 66, 4, 67, 9, 67, 4, 68, 9, 68, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4,
	6, 4, 142, 10, 4, 13, 4, 14, 4, 143, 3, 4, 3, 4, 3, 5, 3, 5, 5, 5, 150,
	10, 5, 3, 6, 3, 6, 5, 6, 154, 10, 6, 3, 7, 3, 7, 5, 7, 158, 10, 7, 3, 7,
	5, 7, 161, 10, 7, 3, 7, 5, 7, 164, 10, 7, 3, 7, 3, 7, 3, 7, 5, 7, 169,
	10, 7, 3, 8, 3, 8, 3, 9, 3, 9, 6, 9, 175, 10, 9, 13, 9, 14, 9, 176, 3,
	9, 3, 9, 3, 10, 3, 10, 3, 10, 5, 10, 184, 10, 10, 3, 11, 5, 11, 187, 10,
	11, 3, 11, 3, 11, 5, 11, 191, 10, 11, 3, 11, 5, 11, 194, 10, 11, 3, 11,
	5, 11, 197, 10, 11, 3, 12, 3, 12, 6, 12, 201, 10, 12, 13, 12, 14, 12, 202,
	3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 15, 3,
	15, 3, 15, 5, 15, 217, 10, 15, 3, 16, 3, 16, 3, 16, 3, 16, 5, 16, 223,
	10, 16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19, 3, 19,
	5, 19, 234, 10, 19, 3, 19, 5, 19, 237, 10, 19, 3, 19, 3, 19, 3, 20, 3,
	20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 3, 20, 5, 20, 250, 10, 20,
	3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 261,
	10, 22, 3, 23, 3, 23, 3, 24, 3, 24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 27,
	3, 27, 3, 27, 3, 27, 6, 27, 275, 10, 27, 13, 27, 14, 27, 276, 3, 27, 3,
	27, 5, 27, 281, 10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 6, 28, 287, 10, 28,
	13, 28, 14, 28, 288, 3, 28, 3, 28, 5, 28, 293, 10, 28, 3, 29, 3, 29, 3,
	29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 6, 31, 304, 10, 31, 13, 31,
	14, 31, 305, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 314, 10,
	32, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 5, 34, 321, 10, 34, 3, 35, 3, 35,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 5,
	37, 335, 10, 37, 3, 38, 6, 38, 338, 10, 38, 13, 38, 14, 38, 339, 3, 39,
	3, 39, 3, 39, 5, 39, 345, 10, 39, 3, 40, 3, 40, 3, 40, 3, 40, 5, 40, 351,
	10, 40, 3, 41, 3, 41, 5, 41, 355, 10, 41, 3, 41, 5, 41, 358, 10, 41, 3,
	41, 3, 41, 6, 41, 362, 10, 41, 13, 41, 14, 41, 363, 3, 41, 3, 41, 3, 42,
	3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 3,
	44, 5, 44, 380, 10, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45,
	388, 10, 45, 3, 46, 5, 46, 391, 10, 46, 3, 46, 3, 46, 3, 46, 5, 46, 396,
	10, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 5, 48, 404, 10, 48, 3,
	48, 3, 48, 3, 48, 5, 48, 409, 10, 48, 3, 48, 5, 48, 412, 10, 48, 3, 48,
	5, 48, 415, 10, 48, 3, 49, 3, 49, 3, 49, 3, 49, 5, 49, 421, 10, 49, 3,
	49, 5, 49, 424, 10, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 5, 49,
	432, 10, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3,
	49, 3, 49, 3, 49, 3, 49, 3, 49, 5, 49, 447, 10, 49, 3, 49, 5, 49, 450,
	10, 49, 5, 49, 452, 10, 49, 3, 50, 3, 50, 3, 50, 5, 50, 457, 10, 50, 3,
	50, 3, 50, 3, 50, 3, 50, 5, 50, 463, 10, 50, 3, 50, 7, 50, 466, 10, 50,
	12, 50, 14, 50, 469, 11, 50, 3, 51, 3, 51, 6, 51, 473, 10, 51, 13, 51,
	14, 51, 474, 3, 51, 3, 51, 3, 52, 5, 52, 480, 10, 52, 3, 52, 3, 52, 5,
	52, 484, 10, 52, 3, 52, 3, 52, 3, 52, 5, 52, 489, 10, 52, 3, 53, 3, 53,
	7, 53, 493, 10, 53, 12, 53, 14, 53, 496, 11, 53, 3, 53, 3, 53, 3, 54, 5,
	54, 501, 10, 54, 3, 54, 3, 54, 3, 54, 3, 54, 5, 54, 507, 10, 54, 3, 54,
	5, 54, 510, 10, 54, 3, 54, 5, 54, 513, 10, 54, 3, 55, 5, 55, 516, 10, 55,
	3, 55, 3, 55, 3, 55, 5, 55, 521, 10, 55, 3, 55, 5, 55, 524, 10, 55, 3,
	56, 3, 56, 3, 56, 3, 56, 5, 56, 530, 10, 56, 3, 56, 3, 56, 3, 56, 3, 56,
	3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 3, 56, 5, 56, 545,
	10, 56, 3, 56, 5, 56, 548, 10, 56, 5, 56, 550, 10, 56, 3, 57, 5, 57, 553,
	10, 57, 3, 57, 3, 57, 3, 57, 5, 57, 558, 10, 57, 3, 57, 5, 57, 561, 10,
	57, 3, 58, 3, 58, 3, 58, 5, 58, 566, 10, 58, 3, 58, 3, 58, 3, 58, 3, 58,
	3, 58, 7, 58, 573, 10, 58, 12, 58, 14, 58, 576, 11, 58, 3, 59, 3, 59, 3,
	59, 3, 59, 5, 59, 582, 10, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59, 3, 59,
	3, 59, 5, 59, 591, 10, 59, 3, 60, 5, 60, 594, 10, 60, 3, 60, 3, 60, 3,
	60, 5, 60, 599, 10, 60, 3, 60, 5, 60, 602, 10, 60, 3, 61, 3, 61, 6, 61,
	606, 10, 61, 13, 61, 14, 61, 607, 3, 61, 3, 61, 3, 62, 5, 62, 613, 10,
	62, 3, 62, 3, 62, 5, 62, 617, 10, 62, 3, 63, 3, 63, 3, 63, 3, 63, 5, 63,
	623, 10, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 3, 63, 5, 63, 632,
	10, 63, 3, 64, 5, 64, 635, 10, 64, 3, 64, 3, 64, 3, 64, 5, 64, 640, 10,
	64, 3, 64, 5, 64, 643, 10, 64, 3, 65, 3, 65, 6, 65, 647, 10, 65, 13, 65,
	14, 65, 648, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 5, 66, 657, 10,
	66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 3, 66, 5, 66, 666, 10, 66,
	3, 67, 5, 67, 669, 10, 67, 3, 67, 3, 67, 3, 67, 3, 67, 5, 67, 675, 10,
	67, 3, 67, 5, 67, 678, 10, 67, 3, 67, 3, 67, 3, 67, 3, 68, 3, 68, 5, 68,
	685, 10, 68, 3, 68, 3, 68, 3, 68, 3, 68, 3, 68, 7, 68, 692, 10, 68, 12,
	68, 14, 68, 695, 11, 68, 3, 68, 2, 5, 98, 114, 134, 69, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
	48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
	84, 86, 88, 90, 92, 94, 96, 98, 100, 102, 104, 106, 108, 110, 112, 114,
	116, 118, 120, 122, 124, 126, 128, 130, 132, 134, 2, 5, 5, 2, 3, 18, 48,
	48, 51, 51, 3, 2, 15, 17, 3, 2, 25, 26, 2, 745, 2, 136, 3, 2, 2, 2, 4,
	138, 3, 2, 2, 2, 6, 141, 3, 2, 2, 2, 8, 149, 3, 2, 2, 2, 10, 153, 3, 2,
	2, 2, 12, 168, 3, 2, 2, 2, 14, 170, 3, 2, 2, 2, 16, 172, 3, 2, 2, 2, 18,
	183, 3, 2, 2, 2, 20, 186, 3, 2, 2, 2, 22, 198, 3, 2, 2, 2, 24, 206, 3,
	2, 2, 2, 26, 210, 3, 2, 2, 2, 28, 213, 3, 2, 2, 2, 30, 218, 3, 2, 2, 2,
	32, 226, 3, 2, 2, 2, 34, 228, 3, 2, 2, 2, 36, 231, 3, 2, 2, 2, 38, 249,
	3, 2, 2, 2, 40, 251, 3, 2, 2, 2, 42, 260, 3, 2, 2, 2, 44, 262, 3, 2, 2,
	2, 46, 264, 3, 2, 2, 2, 48, 266, 3, 2, 2, 2, 50, 268, 3, 2, 2, 2, 52, 280,
	3, 2, 2, 2, 54, 292, 3, 2, 2, 2, 56, 294, 3, 2, 2, 2, 58, 298, 3, 2, 2,
	2, 60, 301, 3, 2, 2, 2, 62, 309, 3, 2, 2, 2, 64, 315, 3, 2, 2, 2, 66, 320,
	3, 2, 2, 2, 68, 322, 3, 2, 2, 2, 70, 324, 3, 2, 2, 2, 72, 334, 3, 2, 2,
	2, 74, 337, 3, 2, 2, 2, 76, 341, 3, 2, 2, 2, 78, 350, 3, 2, 2, 2, 80, 352,
	3, 2, 2, 2, 82, 367, 3, 2, 2, 2, 84, 371, 3, 2, 2, 2, 86, 379, 3, 2, 2,
	2, 88, 387, 3, 2, 2, 2, 90, 390, 3, 2, 2, 2, 92, 397, 3, 2, 2, 2, 94, 403,
	3, 2, 2, 2, 96, 451, 3, 2, 2, 2, 98, 453, 3, 2, 2, 2, 100, 470, 3, 2, 2,
	2, 102, 479, 3, 2, 2, 2, 104, 490, 3, 2, 2, 2, 106, 500, 3, 2, 2, 2, 108,
	515, 3, 2, 2, 2, 110, 549, 3, 2, 2, 2, 112, 552, 3, 2, 2, 2, 114, 562,
	3, 2, 2, 2, 116, 590, 3, 2, 2, 2, 118, 593, 3, 2, 2, 2, 120, 603, 3, 2,
	2, 2, 122, 612, 3, 2, 2, 2, 124, 631, 3, 2, 2, 2, 126, 634, 3, 2, 2, 2,
	128, 644, 3, 2, 2, 2, 130, 665, 3, 2, 2, 2, 132, 668, 3, 2, 2, 2, 134,
	682, 3, 2, 2, 2, 136, 137, 5, 6, 4, 2, 137, 3, 3, 2, 2, 2, 138, 139, 9,
	2, 2, 2, 139, 5, 3, 2, 2, 2, 140, 142, 5, 8, 5, 2, 141, 140, 3, 2, 2, 2,
	142, 143, 3, 2, 2, 2, 143, 141, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144,
	145, 3, 2, 2, 2, 145, 146, 7, 2, 2, 3, 146, 7, 3, 2, 2, 2, 147, 150, 5,
	10, 6, 2, 148, 150, 5, 78, 40, 2, 149, 147, 3, 2, 2, 2, 149, 148, 3, 2,
	2, 2, 150, 9, 3, 2, 2, 2, 151, 154, 5, 12, 7, 2, 152, 154, 5, 30, 16, 2,
	153, 151, 3, 2, 2, 2, 153, 152, 3, 2, 2, 2, 154, 11, 3, 2, 2, 2, 155, 157,
	5, 14, 8, 2, 156, 158, 5, 4, 3, 2, 157, 156, 3, 2, 2, 2, 157, 158, 3, 2,
	2, 2, 158, 160, 3, 2, 2, 2, 159, 161, 5, 60, 31, 2, 160, 159, 3, 2, 2,
	2, 160, 161, 3, 2, 2, 2, 161, 163, 3, 2, 2, 2, 162, 164, 5, 74, 38, 2,
	163, 162, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165, 3, 2, 2, 2, 165,
	166, 5, 16, 9, 2, 166, 169, 3, 2, 2, 2, 167, 169, 5, 16, 9, 2, 168, 155,
	3, 2, 2, 2, 168, 167, 3, 2, 2, 2, 169, 13, 3, 2, 2, 2, 170, 171, 9, 3,
	2, 2, 171, 15, 3, 2, 2, 2, 172, 174, 7, 19, 2, 2, 173, 175, 5, 18, 10,
	2, 174, 173, 3, 2, 2, 2, 175, 176, 3, 2, 2, 2, 176, 174, 3, 2, 2, 2, 176,
	177, 3, 2, 2, 2, 177, 178, 3, 2, 2, 2, 178, 179, 7, 20, 2, 2, 179, 17,
	3, 2, 2, 2, 180, 184, 5, 20, 11, 2, 181, 184, 5, 28, 15, 2, 182, 184, 5,
	36, 19, 2, 183, 180, 3, 2, 2, 2, 183, 181, 3, 2, 2, 2, 183, 182, 3, 2,
	2, 2, 184, 19, 3, 2, 2, 2, 185, 187, 5, 26, 14, 2, 186, 185, 3, 2, 2, 2,
	186, 187, 3, 2, 2, 2, 187, 188, 3, 2, 2, 2, 188, 190, 5, 4, 3, 2, 189,
	191, 5, 22, 12, 2, 190, 189, 3, 2, 2, 2, 190, 191, 3, 2, 2, 2, 191, 193,
	3, 2, 2, 2, 192, 194, 5, 74, 38, 2, 193, 192, 3, 2, 2, 2, 193, 194, 3,
	2, 2, 2, 194, 196, 3, 2, 2, 2, 195, 197, 5, 16, 9, 2, 196, 195, 3, 2, 2,
	2, 196, 197, 3, 2, 2, 2, 197, 21, 3, 2, 2, 2, 198, 200, 7, 21, 2, 2, 199,
	201, 5, 24, 13, 2, 200, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 200,
	3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 204, 3, 2, 2, 2, 204, 205, 7, 22,
	2, 2, 205, 23, 3, 2, 2, 2, 206, 207, 5, 4, 3, 2, 207, 208, 7, 23, 2, 2,
	208, 209, 5, 38, 20, 2, 209, 25, 3, 2, 2, 2, 210, 211, 5, 4, 3, 2, 211,
	212, 7, 23, 2, 2, 212, 27, 3, 2, 2, 2, 213, 214, 7, 24, 2, 2, 214, 216,
	5, 32, 17, 2, 215, 217, 5, 74, 38, 2, 216, 215, 3, 2, 2, 2, 216, 217, 3,
	2, 2, 2, 217, 29, 3, 2, 2, 2, 218, 219, 7, 5, 2, 2, 219, 220, 5, 32, 17,
	2, 220, 222, 5, 34, 18, 2, 221, 223, 5, 74, 38, 2, 222, 221, 3, 2, 2, 2,
	222, 223, 3, 2, 2, 2, 223, 224, 3, 2, 2, 2, 224, 225, 5, 16, 9, 2, 225,
	31, 3, 2, 2, 2, 226, 227, 5, 4, 3, 2, 227, 33, 3, 2, 2, 2, 228, 229, 7,
	4, 2, 2, 229, 230, 5, 68, 35, 2, 230, 35, 3, 2, 2, 2, 231, 233, 7, 24,
	2, 2, 232, 234, 5, 34, 18, 2, 233, 232, 3, 2, 2, 2, 233, 234, 3, 2, 2,
	2, 234, 236, 3, 2, 2, 2, 235, 237, 5, 74, 38, 2, 236, 235, 3, 2, 2, 2,
	236, 237, 3, 2, 2, 2, 237, 238, 3, 2, 2, 2, 238, 239, 5, 16, 9, 2, 239,
	37, 3, 2, 2, 2, 240, 250, 5, 58, 30, 2, 241, 250, 5, 42, 22, 2, 242, 250,
	5, 40, 21, 2, 243, 250, 5, 46, 24, 2, 244, 250, 5, 44, 23, 2, 245, 250,
	5, 48, 25, 2, 246, 250, 5, 50, 26, 2, 247, 250, 5, 52, 27, 2, 248, 250,
	5, 54, 28, 2, 249, 240, 3, 2, 2, 2, 249, 241, 3, 2, 2, 2, 249, 242, 3,
	2, 2, 2, 249, 243, 3, 2, 2, 2, 249, 244, 3, 2, 2, 2, 249, 245, 3, 2, 2,
	2, 249, 246, 3, 2, 2, 2, 249, 247, 3, 2, 2, 2, 249, 248, 3, 2, 2, 2, 250,
	39, 3, 2, 2, 2, 251, 252, 7, 43, 2, 2, 252, 41, 3, 2, 2, 2, 253, 254, 7,
	43, 2, 2, 254, 255, 7, 45, 2, 2, 255, 261, 7, 46, 2, 2, 256, 257, 7, 43,
	2, 2, 257, 261, 7, 46, 2, 2, 258, 259, 7, 43, 2, 2, 259, 261, 7, 45, 2,
	2, 260, 253, 3, 2, 2, 2, 260, 256, 3, 2, 2, 2, 260, 258, 3, 2, 2, 2, 261,
	43, 3, 2, 2, 2, 262, 263, 9, 4, 2, 2, 263, 45, 3, 2, 2, 2, 264, 265, 7,
	47, 2, 2, 265, 47, 3, 2, 2, 2, 266, 267, 7, 27, 2, 2, 267, 49, 3, 2, 2,
	2, 268, 269, 5, 4, 3, 2, 269, 51, 3, 2, 2, 2, 270, 271, 7, 28, 2, 2, 271,
	281, 7, 29, 2, 2, 272, 274, 7, 28, 2, 2, 273, 275, 5, 38, 20, 2, 274, 273,
	3, 2, 2, 2, 275, 276, 3, 2, 2, 2, 276, 274, 3, 2, 2, 2, 276, 277, 3, 2,
	2, 2, 277, 278, 3, 2, 2, 2, 278, 279, 7, 29, 2, 2, 279, 281, 3, 2, 2, 2,
	280, 270, 3, 2, 2, 2, 280, 272, 3, 2, 2, 2, 281, 53, 3, 2, 2, 2, 282, 283,
	7, 19, 2, 2, 283, 293, 7, 20, 2, 2, 284, 286, 7, 19, 2, 2, 285, 287, 5,
	56, 29, 2, 286, 285, 3, 2, 2, 2, 287, 288, 3, 2, 2, 2, 288, 286, 3, 2,
	2, 2, 288, 289, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 291, 7, 20, 2, 2,
	291, 293, 3, 2, 2, 2, 292, 282, 3, 2, 2, 2, 292, 284, 3, 2, 2, 2, 293,
	55, 3, 2, 2, 2, 294, 295, 5, 4, 3, 2, 295, 296, 7, 23, 2, 2, 296, 297,
	5, 38, 20, 2, 297, 57, 3, 2, 2, 2, 298, 299, 7, 30, 2, 2, 299, 300, 5,
	4, 3, 2, 300, 59, 3, 2, 2, 2, 301, 303, 7, 21, 2, 2, 302, 304, 5, 62, 32,
	2, 303, 302, 3, 2, 2, 2, 304, 305, 3, 2, 2, 2, 305, 303, 3, 2, 2, 2, 305,
	306, 3, 2, 2, 2, 306, 307, 3, 2, 2, 2, 307, 308, 7, 22, 2, 2, 308, 61,
	3, 2, 2, 2, 309, 310, 5, 58, 30, 2, 310, 311, 7, 23, 2, 2, 311, 313, 5,
	66, 34, 2, 312, 314, 5, 64, 33, 2, 313, 312, 3, 2, 2, 2, 313, 314, 3, 2,
	2, 2, 314, 63, 3, 2, 2, 2, 315, 316, 5, 38, 20, 2, 316, 65, 3, 2, 2, 2,
	317, 321, 5, 68, 35, 2, 318, 321, 5, 70, 36, 2, 319, 321, 5, 72, 37, 2,
	320, 317, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320, 319, 3, 2, 2, 2, 321,
	67, 3, 2, 2, 2, 322, 323, 5, 4, 3, 2, 323, 69, 3, 2, 2, 2, 324, 325, 7,
	28, 2, 2, 325, 326, 5, 66, 34, 2, 326, 327, 7, 29, 2, 2, 327, 71, 3, 2,
	2, 2, 328, 329, 5, 68, 35, 2, 329, 330, 7, 31, 2, 2, 330, 335, 3, 2, 2,
	2, 331, 332, 5, 70, 36, 2, 332, 333, 7, 31, 2, 2, 333, 335, 3, 2, 2, 2,
	334, 328, 3, 2, 2, 2, 334, 331, 3, 2, 2, 2, 335, 73, 3, 2, 2, 2, 336, 338,
	5, 76, 39, 2, 337, 336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 337, 3,
	2, 2, 2, 339, 340, 3, 2, 2, 2, 340, 75, 3, 2, 2, 2, 341, 342, 7, 32, 2,
	2, 342, 344, 5, 4, 3, 2, 343, 345, 5, 22, 12, 2, 344, 343, 3, 2, 2, 2,
	344, 345, 3, 2, 2, 2, 345, 77, 3, 2, 2, 2, 346, 351, 5, 80, 41, 2, 347,
	351, 5, 86, 44, 2, 348, 351, 5, 88, 45, 2, 349, 351, 5, 132, 67, 2, 350,
	346, 3, 2, 2, 2, 350, 347, 3, 2, 2, 2, 350, 348, 3, 2, 2, 2, 350, 349,
	3, 2, 2, 2, 351, 79, 3, 2, 2, 2, 352, 354, 7, 8, 2, 2, 353, 355, 5, 4,
	3, 2, 354, 353, 3, 2, 2, 2, 354, 355, 3, 2, 2, 2, 355, 357, 3, 2, 2, 2,
	356, 358, 5, 74, 38, 2, 357, 356, 3, 2, 2, 2, 357, 358, 3, 2, 2, 2, 358,
	359, 3, 2, 2, 2, 359, 361, 7, 19, 2, 2, 360, 362, 5, 82, 42, 2, 361, 360,
	3, 2, 2, 2, 362, 363, 3, 2, 2, 2, 363, 361, 3, 2, 2, 2, 363, 364, 3, 2,
	2, 2, 364, 365, 3, 2, 2, 2, 365, 366, 7, 20, 2, 2, 366, 81, 3, 2, 2, 2,
	367, 368, 5, 14, 8, 2, 368, 369, 7, 23, 2, 2, 369, 370, 5, 68, 35, 2, 370,
	83, 3, 2, 2, 2, 371, 372, 5, 46, 24, 2, 372, 85, 3, 2, 2, 2, 373, 380,
	5, 90, 46, 2, 374, 380, 5, 94, 48, 2, 375, 380, 5, 108, 55, 2, 376, 380,
	5, 112, 57, 2, 377, 380, 5, 118, 60, 2, 378, 380, 5, 126, 64, 2, 379, 373,
	3, 2, 2, 2, 379, 374, 3, 2, 2, 2, 379, 375, 3, 2, 2, 2, 379, 376, 3, 2,
	2, 2, 379, 377, 3, 2, 2, 2, 379, 378, 3, 2, 2, 2, 380, 87, 3, 2, 2, 2,
	381, 388, 5, 92, 47, 2, 382, 388, 5, 96, 49, 2, 383, 388, 5, 110, 56, 2,
	384, 388, 5, 116, 59, 2, 385, 388, 5, 124, 63, 2, 386, 388, 5, 130, 66,
	2, 387, 381, 3, 2, 2, 2, 387, 382, 3, 2, 2, 2, 387, 383, 3, 2, 2, 2, 387,
	384, 3, 2, 2, 2, 387, 385, 3, 2, 2, 2, 387, 386, 3, 2, 2, 2, 388, 89, 3,
	2, 2, 2, 389, 391, 5, 84, 43, 2, 390, 389, 3, 2, 2, 2, 390, 391, 3, 2,
	2, 2, 391, 392, 3, 2, 2, 2, 392, 393, 7, 13, 2, 2, 393, 395, 5, 4, 3, 2,
	394, 396, 5, 74, 38, 2, 395, 394, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396,
	91, 3, 2, 2, 2, 397, 398, 7, 6, 2, 2, 398, 399, 7, 13, 2, 2, 399, 400,
	5, 4, 3, 2, 400, 401, 5, 74, 38, 2, 401, 93, 3, 2, 2, 2, 402, 404, 5, 84,
	43, 2, 403, 402, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2,
	405, 406, 7, 3, 2, 2, 406, 408, 5, 4, 3, 2, 407, 409, 5, 98, 50, 2, 408,
	407, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 411, 3, 2, 2, 2, 410, 412,
	5, 74, 38, 2, 411, 410, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 414, 3,
	2, 2, 2, 413, 415, 5, 100, 51, 2, 414, 413, 3, 2, 2, 2, 414, 415, 3, 2,
	2, 2, 415, 95, 3, 2, 2, 2, 416, 417, 7, 6, 2, 2, 417, 418, 7, 3, 2, 2,
	418, 420, 5, 4, 3, 2, 419, 421, 5, 98, 50, 2, 420, 419, 3, 2, 2, 2, 420,
	421, 3, 2, 2, 2, 421, 423, 3, 2, 2, 2, 422, 424, 5, 74, 38, 2, 423, 422,
	3, 2, 2, 2, 423, 424, 3, 2, 2, 2, 424, 425, 3, 2, 2, 2, 425, 426, 5, 100,
	51, 2, 426, 452, 3, 2, 2, 2, 427, 428, 7, 6, 2, 2, 428, 429, 7, 3, 2, 2,
	429, 431, 5, 4, 3, 2, 430, 432, 5, 98, 50, 2, 431, 430, 3, 2, 2, 2, 431,
	432, 3, 2, 2, 2, 432, 433, 3, 2, 2, 2, 433, 434, 5, 74, 38, 2, 434, 452,
	3, 2, 2, 2, 435, 436, 7, 6, 2, 2, 436, 437, 7, 3, 2, 2, 437, 438, 5, 4,
	3, 2, 438, 439, 5, 98, 50, 2, 439, 452, 3, 2, 2, 2, 440, 441, 7, 6, 2,
	2, 441, 442, 7, 3, 2, 2, 442, 443, 5, 4, 3, 2, 443, 444, 7, 18, 2, 2, 444,
	446, 5, 4, 3, 2, 445, 447, 5, 74, 38, 2, 446, 445, 3, 2, 2, 2, 446, 447,
	3, 2, 2, 2, 447, 449, 3, 2, 2, 2, 448, 450, 5, 100, 51, 2, 449, 448, 3,
	2, 2, 2, 449, 450, 3, 2, 2, 2, 450, 452, 3, 2, 2, 2, 451, 416, 3, 2, 2,
	2, 451, 427, 3, 2, 2, 2, 451, 435, 3, 2, 2, 2, 451, 440, 3, 2, 2, 2, 452,
	97, 3, 2, 2, 2, 453, 454, 8, 50, 1, 2, 454, 456, 7, 7, 2, 2, 455, 457,
	7, 33, 2, 2, 456, 455, 3, 2, 2, 2, 456, 457, 3, 2, 2, 2, 457, 458, 3, 2,
	2, 2, 458, 459, 5, 68, 35, 2, 459, 467, 3, 2, 2, 2, 460, 462, 12, 3, 2,
	2, 461, 463, 7, 33, 2, 2, 462, 461, 3, 2, 2, 2, 462, 463, 3, 2, 2, 2, 463,
	464, 3, 2, 2, 2, 464, 466, 5, 68, 35, 2, 465, 460, 3, 2, 2, 2, 466, 469,
	3, 2, 2, 2, 467, 465, 3, 2, 2, 2, 467, 468, 3, 2, 2, 2, 468, 99, 3, 2,
	2, 2, 469, 467, 3, 2, 2, 2, 470, 472, 7, 19, 2, 2, 471, 473, 5, 102, 52,
	2, 472, 471, 3, 2, 2, 2, 473, 474, 3, 2, 2, 2, 474, 472, 3, 2, 2, 2, 474,
	475, 3, 2, 2, 2, 475, 476, 3, 2, 2, 2, 476, 477, 7, 20, 2, 2, 477, 101,
	3, 2, 2, 2, 478, 480, 5, 84, 43, 2, 479, 478, 3, 2, 2, 2, 479, 480, 3,
	2, 2, 2, 480, 481, 3, 2, 2, 2, 481, 483, 5, 4, 3, 2, 482, 484, 5, 104,
	53, 2, 483, 482, 3, 2, 2, 2, 483, 484, 3, 2, 2, 2, 484, 485, 3, 2, 2, 2,
	485, 486, 7, 23, 2, 2, 486, 488, 5, 66, 34, 2, 487, 489, 5, 74, 38, 2,
	488, 487, 3, 2, 2, 2, 488, 489, 3, 2, 2, 2, 489, 103, 3, 2, 2, 2, 490,
	494, 7, 21, 2, 2, 491, 493, 5, 106, 54, 2, 492, 491, 3, 2, 2, 2, 493, 496,
	3, 2, 2, 2, 494, 492, 3, 2, 2, 2, 494, 495, 3, 2, 2, 2, 495, 497, 3, 2,
	2, 2, 496, 494, 3, 2, 2, 2, 497, 498, 7, 22, 2, 2, 498, 105, 3, 2, 2, 2,
	499, 501, 5, 84, 43, 2, 500, 499, 3, 2, 2, 2, 500, 501, 3, 2, 2, 2, 501,
	502, 3, 2, 2, 2, 502, 503, 5, 4, 3, 2, 503, 504, 7, 23, 2, 2, 504, 509,
	5, 66, 34, 2, 505, 507, 7, 34, 2, 2, 506, 505, 3, 2, 2, 2, 506, 507, 3,
	2, 2, 2, 507, 508, 3, 2, 2, 2, 508, 510, 5, 64, 33, 2, 509, 506, 3, 2,
	2, 2, 509, 510, 3, 2, 2, 2, 510, 512, 3, 2, 2, 2, 511, 513, 5, 74, 38,
	2, 512, 511, 3, 2, 2, 2, 512, 513, 3, 2, 2, 2, 513, 107, 3, 2, 2, 2, 514,
	516, 5, 84, 43, 2, 515, 514, 3, 2, 2, 2, 515, 516, 3, 2, 2, 2, 516, 517,
	3, 2, 2, 2, 517, 518, 7, 11, 2, 2, 518, 520, 5, 4, 3, 2, 519, 521, 5, 74,
	38, 2, 520, 519, 3, 2, 2, 2, 520, 521, 3, 2, 2, 2, 521, 523, 3, 2, 2, 2,
	522, 524, 5, 100, 51, 2, 523, 522, 3, 2, 2, 2, 523, 524, 3, 2, 2, 2, 524,
	109, 3, 2, 2, 2, 525, 526, 7, 6, 2, 2, 526, 527, 7, 11, 2, 2, 527, 529,
	5, 4, 3, 2, 528, 530, 5, 74, 38, 2, 529, 528, 3, 2, 2, 2, 529, 530, 3,
	2, 2, 2, 530, 531, 3, 2, 2, 2, 531, 532, 5, 100, 51, 2, 532, 550, 3, 2,
	2, 2, 533, 534, 7, 6, 2, 2, 534, 535, 7, 11, 2, 2, 535, 536, 5, 4, 3, 2,
	536, 537, 5, 74, 38, 2, 537, 550, 3, 2, 2, 2, 538, 539, 7, 6, 2, 2, 539,
	540, 7, 11, 2, 2, 540, 541, 5, 4, 3, 2, 541, 542, 7, 18, 2, 2, 542, 544,
	5, 4, 3, 2, 543, 545, 5, 74, 38, 2, 544, 543, 3, 2, 2, 2, 544, 545, 3,
	2, 2, 2, 545, 547, 3, 2, 2, 2, 546, 548, 5, 100, 51, 2, 547, 546, 3, 2,
	2, 2, 547, 548, 3, 2, 2, 2, 548, 550, 3, 2, 2, 2, 549, 525, 3, 2, 2, 2,
	549, 533, 3, 2, 2, 2, 549, 538, 3, 2, 2, 2, 550, 111, 3, 2, 2, 2, 551,
	553, 5, 84, 43, 2, 552, 551, 3, 2, 2, 2, 552, 553, 3, 2, 2, 2, 553, 554,
	3, 2, 2, 2, 554, 555, 7, 10, 2, 2, 555, 557, 5, 4, 3, 2, 556, 558, 5, 74,
	38, 2, 557, 556, 3, 2, 2, 2, 557, 558, 3, 2, 2, 2, 558, 560, 3, 2, 2, 2,
	559, 561, 5, 114, 58, 2, 560, 559, 3, 2, 2, 2, 560, 561, 3, 2, 2, 2, 561,
	113, 3, 2, 2, 2, 562, 563, 8, 58, 1, 2, 563, 565, 7, 34, 2, 2, 564, 566,
	7, 35, 2, 2, 565, 564, 3, 2, 2, 2, 565, 566, 3, 2, 2, 2, 566, 567, 3, 2,
	2, 2, 567, 568, 5, 68, 35, 2, 568, 574, 3, 2, 2, 2, 569, 570, 12, 3, 2,
	2, 570, 571, 7, 35, 2, 2, 571, 573, 5, 68, 35, 2, 572, 569, 3, 2, 2, 2,
	573, 576, 3, 2, 2, 2, 574, 572, 3, 2, 2, 2, 574, 575, 3, 2, 2, 2, 575,
	115, 3, 2, 2, 2, 576, 574, 3, 2, 2, 2, 577, 578, 7, 6, 2, 2, 578, 579,
	7, 10, 2, 2, 579, 581, 5, 4, 3, 2, 580, 582, 5, 74, 38, 2, 581, 580, 3,
	2, 2, 2, 581, 582, 3, 2, 2, 2, 582, 583, 3, 2, 2, 2, 583, 584, 5, 114,
	58, 2, 584, 591, 3, 2, 2, 2, 585, 586, 7, 6, 2, 2, 586, 587, 7, 10, 2,
	2, 587, 588, 5, 4, 3, 2, 588, 589, 5, 74, 38, 2, 589, 591, 3, 2, 2, 2,
	590, 577, 3, 2, 2, 2, 590, 585, 3, 2, 2, 2, 591, 117, 3, 2, 2, 2, 592,
	594, 5, 84, 43, 2, 593, 592, 3, 2, 2, 2, 593, 594, 3, 2, 2, 2, 594, 595,
	3, 2, 2, 2, 595, 596, 7, 9, 2, 2, 596, 598, 5, 4, 3, 2, 597, 599, 5, 74,
	38, 2, 598, 597, 3, 2, 2, 2, 598, 599, 3, 2, 2, 2, 599, 601, 3, 2, 2, 2,
	600, 602, 5, 120, 61, 2, 601, 600, 3, 2, 2, 2, 601, 602, 3, 2, 2, 2, 602,
	119, 3, 2, 2, 2, 603, 605, 7, 19, 2, 2, 604, 606, 5, 122, 62, 2, 605, 604,
	3, 2, 2, 2, 606, 607, 3, 2, 2, 2, 607, 605, 3, 2, 2, 2, 607, 608, 3, 2,
	2, 2, 608, 609, 3, 2, 2, 2, 609, 610, 7, 20, 2, 2, 610, 121, 3, 2, 2, 2,
	611, 613, 5, 84, 43, 2, 612, 611, 3, 2, 2, 2, 612, 613, 3, 2, 2, 2, 613,
	614, 3, 2, 2, 2, 614, 616, 5, 50, 26, 2, 615, 617, 5, 74, 38, 2, 616, 615,
	3, 2, 2, 2, 616, 617, 3, 2, 2, 2, 617, 123, 3, 2, 2, 2, 618, 619, 7, 6,
	2, 2, 619, 620, 7, 9, 2, 2, 620, 622, 5, 4, 3, 2, 621, 623, 5, 74, 38,
	2, 622, 621, 3, 2, 2, 2, 622, 623, 3, 2, 2, 2, 623, 624, 3, 2, 2, 2, 624,
	625, 5, 120, 61, 2, 625, 632, 3, 2, 2, 2, 626, 627, 7, 6, 2, 2, 627, 628,
	7, 9, 2, 2, 628, 629, 5, 4, 3, 2, 629, 630, 5, 74, 38, 2, 630, 632, 3,
	2, 2, 2, 631, 618, 3, 2, 2, 2, 631, 626, 3, 2, 2, 2, 632, 125, 3, 2, 2,
	2, 633, 635, 5, 84, 43, 2, 634, 633, 3, 2, 2, 2, 634, 635, 3, 2, 2, 2,
	635, 636, 3, 2, 2, 2, 636, 637, 7, 12, 2, 2, 637, 639, 5, 4, 3, 2, 638,
	640, 5, 74, 38, 2, 639, 638, 3, 2, 2, 2, 639, 640, 3, 2, 2, 2, 640, 642,
	3, 2, 2, 2, 641, 643, 5, 128, 65, 2, 642, 641, 3, 2, 2, 2, 642, 643, 3,
	2, 2, 2, 643, 127, 3, 2, 2, 2, 644, 646, 7, 19, 2, 2, 645, 647, 5, 106,
	54, 2, 646, 645, 3, 2, 2, 2, 647, 648, 3, 2, 2, 2, 648, 646, 3, 2, 2, 2,
	648, 649, 3, 2, 2, 2, 649, 650, 3, 2, 2, 2, 650, 651, 7, 20, 2, 2, 651,
	129, 3, 2, 2, 2, 652, 653, 7, 6, 2, 2, 653, 654, 7, 12, 2, 2, 654, 656,
	5, 4, 3, 2, 655, 657, 5, 74, 38, 2, 656, 655, 3, 2, 2, 2, 656, 657, 3,
	2, 2, 2, 657, 658, 3, 2, 2, 2, 658, 659, 5, 128, 65, 2, 659, 666, 3, 2,
	2, 2, 660, 661, 7, 6, 2, 2, 661, 662, 7, 12, 2, 2, 662, 663, 5, 4, 3, 2,
	663, 664, 5, 74, 38, 2, 664, 666, 3, 2, 2, 2, 665, 652, 3, 2, 2, 2, 665,
	660, 3, 2, 2, 2, 666, 131, 3, 2, 2, 2, 667, 669, 5, 84, 43, 2, 668, 667,
	3, 2, 2, 2, 668, 669, 3, 2, 2, 2, 669, 670, 3, 2, 2, 2, 670, 671, 7, 14,
	2, 2, 671, 672, 7, 32, 2, 2, 672, 674, 5, 4, 3, 2, 673, 675, 5, 74, 38,
	2, 674, 673, 3, 2, 2, 2, 674, 675, 3, 2, 2, 2, 675, 677, 3, 2, 2, 2, 676,
	678, 5, 104, 53, 2, 677, 676, 3, 2, 2, 2, 677, 678, 3, 2, 2, 2, 678, 679,
	3, 2, 2, 2, 679, 680, 7, 4, 2, 2, 680, 681, 5, 134, 68, 2, 681, 133, 3,
	2, 2, 2, 682, 684, 8, 68, 1, 2, 683, 685, 7, 35, 2, 2, 684, 683, 3, 2,
	2, 2, 684, 685, 3, 2, 2, 2, 685, 686, 3, 2, 2, 2, 686, 687, 7, 48, 2, 2,
	687, 693, 3, 2, 2, 2, 688, 689, 12, 3, 2, 2, 689, 690, 7, 35, 2, 2, 690,
	692, 7, 48, 2, 2, 691, 688, 3, 2, 2, 2, 692, 695, 3, 2, 2, 2, 693, 691,
	3, 2, 2, 2, 693, 694, 3, 2, 2, 2, 694, 135, 3, 2, 2, 2, 695, 693, 3, 2,
	2, 2, 95, 143, 149, 153, 157, 160, 163, 168, 176, 183, 186, 190, 193, 196,
	202, 216, 222, 233, 236, 249, 260, 276, 280, 288, 292, 305, 313, 320, 334,
	339, 344, 350, 354, 357, 363, 379, 387, 390, 395, 403, 408, 411, 414, 420,
	423, 431, 446, 449, 451, 456, 462, 467, 474, 479, 483, 488, 494, 500, 506,
	509, 512, 515, 520, 523, 529, 544, 547, 549, 552, 557, 560, 565, 574, 581,
	590, 593, 598, 601, 607, 612, 616, 622, 631, 634, 639, 642, 648, 656, 665,
	668, 674, 677, 684, 693,
}
var literalNames = []string{
	"", "'type'", "'on'", "'fragment'", "'extend'", "'implements'", "'schema'",
	"'enum'", "'union'", "'interface'", "'input'", "'scalar'", "'directive'",
	"'query'", "'mutation'", "'subscription'", "'by'", "'{'", "'}'", "'('",
	"')'", "':'", "'...'", "'true'", "'false'", "'null'", "'['", "']'", "'$'",
	"'!'", "'@'", "'&'", "'='", "'|'", "", "", "", "", "", "", "','", "", "'-'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "WS", "Ignored",
	"UnicodeBOM", "WhiteSpace", "LineTerminator", "Comment", "Comma", "IntegerPart",
	"NegativeSign", "FractionalPart", "ExponentPart", "StringConst", "DirectiveLocation",
	"ExecutableDirectiveLocation", "TypeSystemDirectiveLocation", "NAME",
}

var ruleNames = []string{
	"graphql", "name", "document", "definition", "executableDefinition", "operationDefinition",
	"operationType", "selectionSet", "selection", "field", "arguments", "argument",
	"alias", "fragmentSpread", "fragmentDefinition", "fragmentName", "typeCondition",
	"inlineFragment", "value", "intValue", "floatValue", "booleanValue", "stringValue",
	"nullValue", "enumValue", "listValue", "objectValue", "objectField", "variable",
	"variableDefinitions", "variableDefinition", "defaultValue", "typeSpec",
	"namedType", "listType", "nonNullType", "directives", "directive", "typeSystemDefinition",
	"schemaDefinition", "operationTypeDefinition", "description", "typeDefinition",
	"typeExtension", "scalarTypeDefinition", "scalarTypeExtension", "objectTypeDefinition",
	"objectTypeExtension", "implementsInterfaces", "fieldsDefinition", "fieldDefinition",
	"argumentsDefinition", "inputValueDefinition", "interfaceTypeDefinition",
	"interfaceTypeExtension", "unionTypeDefinition", "unionMemberTypes", "unionTypeExtension",
	"enumTypeDefinition", "enumValuesDefinition", "enumValueDefinition", "enumTypeExtension",
	"inputObjectTypeDefinition", "inputFieldsDefinition", "inputObjectTypeExtension",
	"directiveDefinition", "directiveLocations",
}

type GraphQLParser struct {
	*antlr.BaseParser
}

// NewGraphQLParser produces a new parser instance for the optional input antlr.TokenStream.
//
// The *GraphQLParser instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewGraphQLParser(input antlr.TokenStream) *GraphQLParser {
	this := new(GraphQLParser)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "GraphQL.g4"

	return this
}

// GraphQLParser tokens.
const (
	GraphQLParserEOF                         = antlr.TokenEOF
	GraphQLParserT__0                        = 1
	GraphQLParserT__1                        = 2
	GraphQLParserT__2                        = 3
	GraphQLParserT__3                        = 4
	GraphQLParserT__4                        = 5
	GraphQLParserT__5                        = 6
	GraphQLParserT__6                        = 7
	GraphQLParserT__7                        = 8
	GraphQLParserT__8                        = 9
	GraphQLParserT__9                        = 10
	GraphQLParserT__10                       = 11
	GraphQLParserT__11                       = 12
	GraphQLParserT__12                       = 13
	GraphQLParserT__13                       = 14
	GraphQLParserT__14                       = 15
	GraphQLParserT__15                       = 16
	GraphQLParserT__16                       = 17
	GraphQLParserT__17                       = 18
	GraphQLParserT__18                       = 19
	GraphQLParserT__19                       = 20
	GraphQLParserT__20                       = 21
	GraphQLParserT__21                       = 22
	GraphQLParserT__22                       = 23
	GraphQLParserT__23                       = 24
	GraphQLParserT__24                       = 25
	GraphQLParserT__25                       = 26
	GraphQLParserT__26                       = 27
	GraphQLParserT__27                       = 28
	GraphQLParserT__28                       = 29
	GraphQLParserT__29                       = 30
	GraphQLParserT__30                       = 31
	GraphQLParserT__31                       = 32
	GraphQLParserT__32                       = 33
	GraphQLParserWS                          = 34
	GraphQLParserIgnored                     = 35
	GraphQLParserUnicodeBOM                  = 36
	GraphQLParserWhiteSpace                  = 37
	GraphQLParserLineTerminator              = 38
	GraphQLParserComment                     = 39
	GraphQLParserComma                       = 40
	GraphQLParserIntegerPart                 = 41
	GraphQLParserNegativeSign                = 42
	GraphQLParserFractionalPart              = 43
	GraphQLParserExponentPart                = 44
	GraphQLParserStringConst                 = 45
	GraphQLParserDirectiveLocation           = 46
	GraphQLParserExecutableDirectiveLocation = 47
	GraphQLParserTypeSystemDirectiveLocation = 48
	GraphQLParserNAME                        = 49
)

// GraphQLParser rules.
const (
	GraphQLParserRULE_graphql                   = 0
	GraphQLParserRULE_name                      = 1
	GraphQLParserRULE_document                  = 2
	GraphQLParserRULE_definition                = 3
	GraphQLParserRULE_executableDefinition      = 4
	GraphQLParserRULE_operationDefinition       = 5
	GraphQLParserRULE_operationType             = 6
	GraphQLParserRULE_selectionSet              = 7
	GraphQLParserRULE_selection                 = 8
	GraphQLParserRULE_field                     = 9
	GraphQLParserRULE_arguments                 = 10
	GraphQLParserRULE_argument                  = 11
	GraphQLParserRULE_alias                     = 12
	GraphQLParserRULE_fragmentSpread            = 13
	GraphQLParserRULE_fragmentDefinition        = 14
	GraphQLParserRULE_fragmentName              = 15
	GraphQLParserRULE_typeCondition             = 16
	GraphQLParserRULE_inlineFragment            = 17
	GraphQLParserRULE_value                     = 18
	GraphQLParserRULE_intValue                  = 19
	GraphQLParserRULE_floatValue                = 20
	GraphQLParserRULE_booleanValue              = 21
	GraphQLParserRULE_stringValue               = 22
	GraphQLParserRULE_nullValue                 = 23
	GraphQLParserRULE_enumValue                 = 24
	GraphQLParserRULE_listValue                 = 25
	GraphQLParserRULE_objectValue               = 26
	GraphQLParserRULE_objectField               = 27
	GraphQLParserRULE_variable                  = 28
	GraphQLParserRULE_variableDefinitions       = 29
	GraphQLParserRULE_variableDefinition        = 30
	GraphQLParserRULE_defaultValue              = 31
	GraphQLParserRULE_typeSpec                  = 32
	GraphQLParserRULE_namedType                 = 33
	GraphQLParserRULE_listType                  = 34
	GraphQLParserRULE_nonNullType               = 35
	GraphQLParserRULE_directives                = 36
	GraphQLParserRULE_directive                 = 37
	GraphQLParserRULE_typeSystemDefinition      = 38
	GraphQLParserRULE_schemaDefinition          = 39
	GraphQLParserRULE_operationTypeDefinition   = 40
	GraphQLParserRULE_description               = 41
	GraphQLParserRULE_typeDefinition            = 42
	GraphQLParserRULE_typeExtension             = 43
	GraphQLParserRULE_scalarTypeDefinition      = 44
	GraphQLParserRULE_scalarTypeExtension       = 45
	GraphQLParserRULE_objectTypeDefinition      = 46
	GraphQLParserRULE_objectTypeExtension       = 47
	GraphQLParserRULE_implementsInterfaces      = 48
	GraphQLParserRULE_fieldsDefinition          = 49
	GraphQLParserRULE_fieldDefinition           = 50
	GraphQLParserRULE_argumentsDefinition       = 51
	GraphQLParserRULE_inputValueDefinition      = 52
	GraphQLParserRULE_interfaceTypeDefinition   = 53
	GraphQLParserRULE_interfaceTypeExtension    = 54
	GraphQLParserRULE_unionTypeDefinition       = 55
	GraphQLParserRULE_unionMemberTypes          = 56
	GraphQLParserRULE_unionTypeExtension        = 57
	GraphQLParserRULE_enumTypeDefinition        = 58
	GraphQLParserRULE_enumValuesDefinition      = 59
	GraphQLParserRULE_enumValueDefinition       = 60
	GraphQLParserRULE_enumTypeExtension         = 61
	GraphQLParserRULE_inputObjectTypeDefinition = 62
	GraphQLParserRULE_inputFieldsDefinition     = 63
	GraphQLParserRULE_inputObjectTypeExtension  = 64
	GraphQLParserRULE_directiveDefinition       = 65
	GraphQLParserRULE_directiveLocations        = 66
)

// IGraphqlContext is an interface to support dynamic dispatch.
type IGraphqlContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGraphqlContext differentiates from other interfaces.
	IsGraphqlContext()
}

type GraphqlContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphqlContext() *GraphqlContext {
	var p = new(GraphqlContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_graphql
	return p
}

func (*GraphqlContext) IsGraphqlContext() {}

func NewGraphqlContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphqlContext {
	var p = new(GraphqlContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_graphql

	return p
}

func (s *GraphqlContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphqlContext) Document() IDocumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDocumentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDocumentContext)
}

func (s *GraphqlContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphqlContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphqlContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterGraphql(s)
	}
}

func (s *GraphqlContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitGraphql(s)
	}
}

func (s *GraphqlContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitGraphql(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Graphql() (localctx IGraphqlContext) {
	localctx = NewGraphqlContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, GraphQLParserRULE_graphql)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(134)
		p.Document()
	}

	return localctx
}

// INameContext is an interface to support dynamic dispatch.
type INameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNameContext differentiates from other interfaces.
	IsNameContext()
}

type NameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNameContext() *NameContext {
	var p = new(NameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_name
	return p
}

func (*NameContext) IsNameContext() {}

func NewNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NameContext {
	var p = new(NameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_name

	return p
}

func (s *NameContext) GetParser() antlr.Parser { return s.parser }

func (s *NameContext) NAME() antlr.TerminalNode {
	return s.GetToken(GraphQLParserNAME, 0)
}

func (s *NameContext) DirectiveLocation() antlr.TerminalNode {
	return s.GetToken(GraphQLParserDirectiveLocation, 0)
}

func (s *NameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterName(s)
	}
}

func (s *NameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitName(s)
	}
}

func (s *NameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Name() (localctx INameContext) {
	localctx = NewNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, GraphQLParserRULE_name)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(136)
		_la = p.GetTokenStream().LA(1)

		if !((((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15))) != 0) || _la == GraphQLParserDirectiveLocation || _la == GraphQLParserNAME) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDocumentContext is an interface to support dynamic dispatch.
type IDocumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDocumentContext differentiates from other interfaces.
	IsDocumentContext()
}

type DocumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDocumentContext() *DocumentContext {
	var p = new(DocumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_document
	return p
}

func (*DocumentContext) IsDocumentContext() {}

func NewDocumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DocumentContext {
	var p = new(DocumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_document

	return p
}

func (s *DocumentContext) GetParser() antlr.Parser { return s.parser }

func (s *DocumentContext) EOF() antlr.TerminalNode {
	return s.GetToken(GraphQLParserEOF, 0)
}

func (s *DocumentContext) AllDefinition() []IDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefinitionContext)(nil)).Elem())
	var tst = make([]IDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefinitionContext)
		}
	}

	return tst
}

func (s *DocumentContext) Definition(i int) IDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefinitionContext)
}

func (s *DocumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DocumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DocumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDocument(s)
	}
}

func (s *DocumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDocument(s)
	}
}

func (s *DocumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDocument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Document() (localctx IDocumentContext) {
	localctx = NewDocumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, GraphQLParserRULE_document)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(139)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__16))) != 0) || _la == GraphQLParserStringConst {
		{
			p.SetState(138)
			p.Definition()
		}

		p.SetState(141)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(143)
		p.Match(GraphQLParserEOF)
	}

	return localctx
}

// IDefinitionContext is an interface to support dynamic dispatch.
type IDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefinitionContext differentiates from other interfaces.
	IsDefinitionContext()
}

type DefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefinitionContext() *DefinitionContext {
	var p = new(DefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_definition
	return p
}

func (*DefinitionContext) IsDefinitionContext() {}

func NewDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefinitionContext {
	var p = new(DefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_definition

	return p
}

func (s *DefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DefinitionContext) ExecutableDefinition() IExecutableDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExecutableDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExecutableDefinitionContext)
}

func (s *DefinitionContext) TypeSystemDefinition() ITypeSystemDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSystemDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSystemDefinitionContext)
}

func (s *DefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDefinition(s)
	}
}

func (s *DefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDefinition(s)
	}
}

func (s *DefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Definition() (localctx IDefinitionContext) {
	localctx = NewDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, GraphQLParserRULE_definition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(147)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__2, GraphQLParserT__12, GraphQLParserT__13, GraphQLParserT__14, GraphQLParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(145)
			p.ExecutableDefinition()
		}

	case GraphQLParserT__0, GraphQLParserT__3, GraphQLParserT__5, GraphQLParserT__6, GraphQLParserT__7, GraphQLParserT__8, GraphQLParserT__9, GraphQLParserT__10, GraphQLParserT__11, GraphQLParserStringConst:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(146)
			p.TypeSystemDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IExecutableDefinitionContext is an interface to support dynamic dispatch.
type IExecutableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExecutableDefinitionContext differentiates from other interfaces.
	IsExecutableDefinitionContext()
}

type ExecutableDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExecutableDefinitionContext() *ExecutableDefinitionContext {
	var p = new(ExecutableDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_executableDefinition
	return p
}

func (*ExecutableDefinitionContext) IsExecutableDefinitionContext() {}

func NewExecutableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExecutableDefinitionContext {
	var p = new(ExecutableDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_executableDefinition

	return p
}

func (s *ExecutableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExecutableDefinitionContext) OperationDefinition() IOperationDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationDefinitionContext)
}

func (s *ExecutableDefinitionContext) FragmentDefinition() IFragmentDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentDefinitionContext)
}

func (s *ExecutableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExecutableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExecutableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterExecutableDefinition(s)
	}
}

func (s *ExecutableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitExecutableDefinition(s)
	}
}

func (s *ExecutableDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitExecutableDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ExecutableDefinition() (localctx IExecutableDefinitionContext) {
	localctx = NewExecutableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, GraphQLParserRULE_executableDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(151)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__12, GraphQLParserT__13, GraphQLParserT__14, GraphQLParserT__16:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(149)
			p.OperationDefinition()
		}

	case GraphQLParserT__2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(150)
			p.FragmentDefinition()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationDefinitionContext is an interface to support dynamic dispatch.
type IOperationDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationDefinitionContext differentiates from other interfaces.
	IsOperationDefinitionContext()
}

type OperationDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationDefinitionContext() *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationDefinition
	return p
}

func (*OperationDefinitionContext) IsOperationDefinitionContext() {}

func NewOperationDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationDefinitionContext {
	var p = new(OperationDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationDefinition

	return p
}

func (s *OperationDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *OperationDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *OperationDefinitionContext) VariableDefinitions() IVariableDefinitionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDefinitionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionsContext)
}

func (s *OperationDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *OperationDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationDefinition(s)
	}
}

func (s *OperationDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationDefinition(s)
	}
}

func (s *OperationDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitOperationDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) OperationDefinition() (localctx IOperationDefinitionContext) {
	localctx = NewOperationDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, GraphQLParserRULE_operationDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(166)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__12, GraphQLParserT__13, GraphQLParserT__14:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(153)
			p.OperationType()
		}
		p.SetState(155)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15))) != 0) || _la == GraphQLParserDirectiveLocation || _la == GraphQLParserNAME {
			{
				p.SetState(154)
				p.Name()
			}

		}
		p.SetState(158)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__18 {
			{
				p.SetState(157)
				p.VariableDefinitions()
			}

		}
		p.SetState(161)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__29 {
			{
				p.SetState(160)
				p.Directives()
			}

		}
		{
			p.SetState(163)
			p.SelectionSet()
		}

	case GraphQLParserT__16:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(165)
			p.SelectionSet()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOperationTypeContext is an interface to support dynamic dispatch.
type IOperationTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationTypeContext differentiates from other interfaces.
	IsOperationTypeContext()
}

type OperationTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeContext() *OperationTypeContext {
	var p = new(OperationTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationType
	return p
}

func (*OperationTypeContext) IsOperationTypeContext() {}

func NewOperationTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeContext {
	var p = new(OperationTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationType

	return p
}

func (s *OperationTypeContext) GetParser() antlr.Parser { return s.parser }
func (s *OperationTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationType(s)
	}
}

func (s *OperationTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationType(s)
	}
}

func (s *OperationTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitOperationType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) OperationType() (localctx IOperationTypeContext) {
	localctx = NewOperationTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, GraphQLParserRULE_operationType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(168)
		_la = p.GetTokenStream().LA(1)

		if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// ISelectionSetContext is an interface to support dynamic dispatch.
type ISelectionSetContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionSetContext differentiates from other interfaces.
	IsSelectionSetContext()
}

type SelectionSetContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionSetContext() *SelectionSetContext {
	var p = new(SelectionSetContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_selectionSet
	return p
}

func (*SelectionSetContext) IsSelectionSetContext() {}

func NewSelectionSetContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionSetContext {
	var p = new(SelectionSetContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_selectionSet

	return p
}

func (s *SelectionSetContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionSetContext) AllSelection() []ISelectionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectionContext)(nil)).Elem())
	var tst = make([]ISelectionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectionContext)
		}
	}

	return tst
}

func (s *SelectionSetContext) Selection(i int) ISelectionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectionContext)
}

func (s *SelectionSetContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionSetContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionSetContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSelectionSet(s)
	}
}

func (s *SelectionSetContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSelectionSet(s)
	}
}

func (s *SelectionSetContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitSelectionSet(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) SelectionSet() (localctx ISelectionSetContext) {
	localctx = NewSelectionSetContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, GraphQLParserRULE_selectionSet)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(170)
		p.Match(GraphQLParserT__16)
	}
	p.SetState(172)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15)|(1<<GraphQLParserT__21))) != 0) || _la == GraphQLParserDirectiveLocation || _la == GraphQLParserNAME {
		{
			p.SetState(171)
			p.Selection()
		}

		p.SetState(174)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(GraphQLParserT__17)
	}

	return localctx
}

// ISelectionContext is an interface to support dynamic dispatch.
type ISelectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectionContext differentiates from other interfaces.
	IsSelectionContext()
}

type SelectionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectionContext() *SelectionContext {
	var p = new(SelectionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_selection
	return p
}

func (*SelectionContext) IsSelectionContext() {}

func NewSelectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectionContext {
	var p = new(SelectionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_selection

	return p
}

func (s *SelectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectionContext) Field() IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *SelectionContext) FragmentSpread() IFragmentSpreadContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentSpreadContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentSpreadContext)
}

func (s *SelectionContext) InlineFragment() IInlineFragmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInlineFragmentContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInlineFragmentContext)
}

func (s *SelectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSelection(s)
	}
}

func (s *SelectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSelection(s)
	}
}

func (s *SelectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitSelection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Selection() (localctx ISelectionContext) {
	localctx = NewSelectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, GraphQLParserRULE_selection)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
			p.Field()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(179)
			p.FragmentSpread()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(180)
			p.InlineFragment()
		}

	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldContext) Alias() IAliasContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAliasContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAliasContext)
}

func (s *FieldContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *FieldContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, GraphQLParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(184)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(183)
			p.Alias()
		}

	}
	{
		p.SetState(186)
		p.Name()
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__18 {
		{
			p.SetState(187)
			p.Arguments()
		}

	}
	p.SetState(191)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(190)
			p.Directives()
		}

	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__16 {
		{
			p.SetState(193)
			p.SelectionSet()
		}

	}

	return localctx
}

// IArgumentsContext is an interface to support dynamic dispatch.
type IArgumentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsContext differentiates from other interfaces.
	IsArgumentsContext()
}

type ArgumentsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsContext() *ArgumentsContext {
	var p = new(ArgumentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_arguments
	return p
}

func (*ArgumentsContext) IsArgumentsContext() {}

func NewArgumentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsContext {
	var p = new(ArgumentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_arguments

	return p
}

func (s *ArgumentsContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsContext) AllArgument() []IArgumentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IArgumentContext)(nil)).Elem())
	var tst = make([]IArgumentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IArgumentContext)
		}
	}

	return tst
}

func (s *ArgumentsContext) Argument(i int) IArgumentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IArgumentContext)
}

func (s *ArgumentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArguments(s)
	}
}

func (s *ArgumentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArguments(s)
	}
}

func (s *ArgumentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitArguments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Arguments() (localctx IArgumentsContext) {
	localctx = NewArgumentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, GraphQLParserRULE_arguments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(196)
		p.Match(GraphQLParserT__18)
	}
	p.SetState(198)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15))) != 0) || _la == GraphQLParserDirectiveLocation || _la == GraphQLParserNAME {
		{
			p.SetState(197)
			p.Argument()
		}

		p.SetState(200)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(202)
		p.Match(GraphQLParserT__19)
	}

	return localctx
}

// IArgumentContext is an interface to support dynamic dispatch.
type IArgumentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentContext differentiates from other interfaces.
	IsArgumentContext()
}

type ArgumentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentContext() *ArgumentContext {
	var p = new(ArgumentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_argument
	return p
}

func (*ArgumentContext) IsArgumentContext() {}

func NewArgumentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentContext {
	var p = new(ArgumentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_argument

	return p
}

func (s *ArgumentContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ArgumentContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ArgumentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArgument(s)
	}
}

func (s *ArgumentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArgument(s)
	}
}

func (s *ArgumentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitArgument(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Argument() (localctx IArgumentContext) {
	localctx = NewArgumentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, GraphQLParserRULE_argument)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Name()
	}
	{
		p.SetState(205)
		p.Match(GraphQLParserT__20)
	}
	{
		p.SetState(206)
		p.Value()
	}

	return localctx
}

// IAliasContext is an interface to support dynamic dispatch.
type IAliasContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAliasContext differentiates from other interfaces.
	IsAliasContext()
}

type AliasContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAliasContext() *AliasContext {
	var p = new(AliasContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_alias
	return p
}

func (*AliasContext) IsAliasContext() {}

func NewAliasContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AliasContext {
	var p = new(AliasContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_alias

	return p
}

func (s *AliasContext) GetParser() antlr.Parser { return s.parser }

func (s *AliasContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *AliasContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AliasContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AliasContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterAlias(s)
	}
}

func (s *AliasContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitAlias(s)
	}
}

func (s *AliasContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitAlias(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Alias() (localctx IAliasContext) {
	localctx = NewAliasContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, GraphQLParserRULE_alias)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(208)
		p.Name()
	}
	{
		p.SetState(209)
		p.Match(GraphQLParserT__20)
	}

	return localctx
}

// IFragmentSpreadContext is an interface to support dynamic dispatch.
type IFragmentSpreadContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentSpreadContext differentiates from other interfaces.
	IsFragmentSpreadContext()
}

type FragmentSpreadContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentSpreadContext() *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentSpread
	return p
}

func (*FragmentSpreadContext) IsFragmentSpreadContext() {}

func NewFragmentSpreadContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentSpreadContext {
	var p = new(FragmentSpreadContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentSpread

	return p
}

func (s *FragmentSpreadContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentSpreadContext) FragmentName() IFragmentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentSpreadContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentSpreadContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentSpreadContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentSpreadContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentSpread(s)
	}
}

func (s *FragmentSpreadContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentSpread(s)
	}
}

func (s *FragmentSpreadContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFragmentSpread(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FragmentSpread() (localctx IFragmentSpreadContext) {
	localctx = NewFragmentSpreadContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, GraphQLParserRULE_fragmentSpread)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(211)
		p.Match(GraphQLParserT__21)
	}
	{
		p.SetState(212)
		p.FragmentName()
	}
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(213)
			p.Directives()
		}

	}

	return localctx
}

// IFragmentDefinitionContext is an interface to support dynamic dispatch.
type IFragmentDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentDefinitionContext differentiates from other interfaces.
	IsFragmentDefinitionContext()
}

type FragmentDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentDefinitionContext() *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition
	return p
}

func (*FragmentDefinitionContext) IsFragmentDefinitionContext() {}

func NewFragmentDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentDefinitionContext {
	var p = new(FragmentDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentDefinition

	return p
}

func (s *FragmentDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentDefinitionContext) FragmentName() IFragmentNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFragmentNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFragmentNameContext)
}

func (s *FragmentDefinitionContext) TypeCondition() ITypeConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *FragmentDefinitionContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *FragmentDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FragmentDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentDefinition(s)
	}
}

func (s *FragmentDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentDefinition(s)
	}
}

func (s *FragmentDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFragmentDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FragmentDefinition() (localctx IFragmentDefinitionContext) {
	localctx = NewFragmentDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, GraphQLParserRULE_fragmentDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(216)
		p.Match(GraphQLParserT__2)
	}
	{
		p.SetState(217)
		p.FragmentName()
	}
	{
		p.SetState(218)
		p.TypeCondition()
	}
	p.SetState(220)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(219)
			p.Directives()
		}

	}
	{
		p.SetState(222)
		p.SelectionSet()
	}

	return localctx
}

// IFragmentNameContext is an interface to support dynamic dispatch.
type IFragmentNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFragmentNameContext differentiates from other interfaces.
	IsFragmentNameContext()
}

type FragmentNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFragmentNameContext() *FragmentNameContext {
	var p = new(FragmentNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fragmentName
	return p
}

func (*FragmentNameContext) IsFragmentNameContext() {}

func NewFragmentNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FragmentNameContext {
	var p = new(FragmentNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fragmentName

	return p
}

func (s *FragmentNameContext) GetParser() antlr.Parser { return s.parser }

func (s *FragmentNameContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FragmentNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FragmentNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FragmentNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFragmentName(s)
	}
}

func (s *FragmentNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFragmentName(s)
	}
}

func (s *FragmentNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFragmentName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FragmentName() (localctx IFragmentNameContext) {
	localctx = NewFragmentNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, GraphQLParserRULE_fragmentName)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(224)
		p.Name()
	}

	return localctx
}

// ITypeConditionContext is an interface to support dynamic dispatch.
type ITypeConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeConditionContext differentiates from other interfaces.
	IsTypeConditionContext()
}

type TypeConditionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeConditionContext() *TypeConditionContext {
	var p = new(TypeConditionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeCondition
	return p
}

func (*TypeConditionContext) IsTypeConditionContext() {}

func NewTypeConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeConditionContext {
	var p = new(TypeConditionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeCondition

	return p
}

func (s *TypeConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeConditionContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *TypeConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeCondition(s)
	}
}

func (s *TypeConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeCondition(s)
	}
}

func (s *TypeConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeCondition() (localctx ITypeConditionContext) {
	localctx = NewTypeConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, GraphQLParserRULE_typeCondition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(226)
		p.Match(GraphQLParserT__1)
	}
	{
		p.SetState(227)
		p.NamedType()
	}

	return localctx
}

// IInlineFragmentContext is an interface to support dynamic dispatch.
type IInlineFragmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInlineFragmentContext differentiates from other interfaces.
	IsInlineFragmentContext()
}

type InlineFragmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInlineFragmentContext() *InlineFragmentContext {
	var p = new(InlineFragmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inlineFragment
	return p
}

func (*InlineFragmentContext) IsInlineFragmentContext() {}

func NewInlineFragmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InlineFragmentContext {
	var p = new(InlineFragmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inlineFragment

	return p
}

func (s *InlineFragmentContext) GetParser() antlr.Parser { return s.parser }

func (s *InlineFragmentContext) SelectionSet() ISelectionSetContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectionSetContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectionSetContext)
}

func (s *InlineFragmentContext) TypeCondition() ITypeConditionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeConditionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeConditionContext)
}

func (s *InlineFragmentContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InlineFragmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InlineFragmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InlineFragmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInlineFragment(s)
	}
}

func (s *InlineFragmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInlineFragment(s)
	}
}

func (s *InlineFragmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInlineFragment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InlineFragment() (localctx IInlineFragmentContext) {
	localctx = NewInlineFragmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, GraphQLParserRULE_inlineFragment)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(229)
		p.Match(GraphQLParserT__21)
	}
	p.SetState(231)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__1 {
		{
			p.SetState(230)
			p.TypeCondition()
		}

	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(233)
			p.Directives()
		}

	}
	{
		p.SetState(236)
		p.SelectionSet()
	}

	return localctx
}

// IValueContext is an interface to support dynamic dispatch.
type IValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsValueContext differentiates from other interfaces.
	IsValueContext()
}

type ValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueContext() *ValueContext {
	var p = new(ValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_value
	return p
}

func (*ValueContext) IsValueContext() {}

func NewValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueContext {
	var p = new(ValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_value

	return p
}

func (s *ValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *ValueContext) FloatValue() IFloatValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFloatValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFloatValueContext)
}

func (s *ValueContext) IntValue() IIntValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntValueContext)
}

func (s *ValueContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *ValueContext) BooleanValue() IBooleanValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanValueContext)
}

func (s *ValueContext) NullValue() INullValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INullValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INullValueContext)
}

func (s *ValueContext) EnumValue() IEnumValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *ValueContext) ListValue() IListValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListValueContext)
}

func (s *ValueContext) ObjectValue() IObjectValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectValueContext)
}

func (s *ValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterValue(s)
	}
}

func (s *ValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitValue(s)
	}
}

func (s *ValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Value() (localctx IValueContext) {
	localctx = NewValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, GraphQLParserRULE_value)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(247)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(238)
			p.Variable()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(239)
			p.FloatValue()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(240)
			p.IntValue()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(241)
			p.StringValue()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(242)
			p.BooleanValue()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(243)
			p.NullValue()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(244)
			p.EnumValue()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(245)
			p.ListValue()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(246)
			p.ObjectValue()
		}

	}

	return localctx
}

// IIntValueContext is an interface to support dynamic dispatch.
type IIntValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntValueContext differentiates from other interfaces.
	IsIntValueContext()
}

type IntValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntValueContext() *IntValueContext {
	var p = new(IntValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_intValue
	return p
}

func (*IntValueContext) IsIntValueContext() {}

func NewIntValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntValueContext {
	var p = new(IntValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_intValue

	return p
}

func (s *IntValueContext) GetParser() antlr.Parser { return s.parser }

func (s *IntValueContext) IntegerPart() antlr.TerminalNode {
	return s.GetToken(GraphQLParserIntegerPart, 0)
}

func (s *IntValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterIntValue(s)
	}
}

func (s *IntValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitIntValue(s)
	}
}

func (s *IntValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitIntValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) IntValue() (localctx IIntValueContext) {
	localctx = NewIntValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, GraphQLParserRULE_intValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(249)
		p.Match(GraphQLParserIntegerPart)
	}

	return localctx
}

// IFloatValueContext is an interface to support dynamic dispatch.
type IFloatValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFloatValueContext differentiates from other interfaces.
	IsFloatValueContext()
}

type FloatValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFloatValueContext() *FloatValueContext {
	var p = new(FloatValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_floatValue
	return p
}

func (*FloatValueContext) IsFloatValueContext() {}

func NewFloatValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FloatValueContext {
	var p = new(FloatValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_floatValue

	return p
}

func (s *FloatValueContext) GetParser() antlr.Parser { return s.parser }

func (s *FloatValueContext) IntegerPart() antlr.TerminalNode {
	return s.GetToken(GraphQLParserIntegerPart, 0)
}

func (s *FloatValueContext) FractionalPart() antlr.TerminalNode {
	return s.GetToken(GraphQLParserFractionalPart, 0)
}

func (s *FloatValueContext) ExponentPart() antlr.TerminalNode {
	return s.GetToken(GraphQLParserExponentPart, 0)
}

func (s *FloatValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FloatValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FloatValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFloatValue(s)
	}
}

func (s *FloatValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFloatValue(s)
	}
}

func (s *FloatValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFloatValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FloatValue() (localctx IFloatValueContext) {
	localctx = NewFloatValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, GraphQLParserRULE_floatValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(258)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(251)
			p.Match(GraphQLParserIntegerPart)
		}
		{
			p.SetState(252)
			p.Match(GraphQLParserFractionalPart)
		}
		{
			p.SetState(253)
			p.Match(GraphQLParserExponentPart)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
			p.Match(GraphQLParserIntegerPart)
		}
		{
			p.SetState(255)
			p.Match(GraphQLParserExponentPart)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(256)
			p.Match(GraphQLParserIntegerPart)
		}
		{
			p.SetState(257)
			p.Match(GraphQLParserFractionalPart)
		}

	}

	return localctx
}

// IBooleanValueContext is an interface to support dynamic dispatch.
type IBooleanValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanValueContext differentiates from other interfaces.
	IsBooleanValueContext()
}

type BooleanValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanValueContext() *BooleanValueContext {
	var p = new(BooleanValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_booleanValue
	return p
}

func (*BooleanValueContext) IsBooleanValueContext() {}

func NewBooleanValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanValueContext {
	var p = new(BooleanValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_booleanValue

	return p
}

func (s *BooleanValueContext) GetParser() antlr.Parser { return s.parser }
func (s *BooleanValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterBooleanValue(s)
	}
}

func (s *BooleanValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitBooleanValue(s)
	}
}

func (s *BooleanValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitBooleanValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) BooleanValue() (localctx IBooleanValueContext) {
	localctx = NewBooleanValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, GraphQLParserRULE_booleanValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(260)
		_la = p.GetTokenStream().LA(1)

		if !(_la == GraphQLParserT__22 || _la == GraphQLParserT__23) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IStringValueContext is an interface to support dynamic dispatch.
type IStringValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringValueContext differentiates from other interfaces.
	IsStringValueContext()
}

type StringValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringValueContext() *StringValueContext {
	var p = new(StringValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_stringValue
	return p
}

func (*StringValueContext) IsStringValueContext() {}

func NewStringValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringValueContext {
	var p = new(StringValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_stringValue

	return p
}

func (s *StringValueContext) GetParser() antlr.Parser { return s.parser }

func (s *StringValueContext) StringConst() antlr.TerminalNode {
	return s.GetToken(GraphQLParserStringConst, 0)
}

func (s *StringValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterStringValue(s)
	}
}

func (s *StringValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitStringValue(s)
	}
}

func (s *StringValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitStringValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) StringValue() (localctx IStringValueContext) {
	localctx = NewStringValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, GraphQLParserRULE_stringValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(GraphQLParserStringConst)
	}

	return localctx
}

// INullValueContext is an interface to support dynamic dispatch.
type INullValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNullValueContext differentiates from other interfaces.
	IsNullValueContext()
}

type NullValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNullValueContext() *NullValueContext {
	var p = new(NullValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_nullValue
	return p
}

func (*NullValueContext) IsNullValueContext() {}

func NewNullValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NullValueContext {
	var p = new(NullValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_nullValue

	return p
}

func (s *NullValueContext) GetParser() antlr.Parser { return s.parser }
func (s *NullValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NullValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterNullValue(s)
	}
}

func (s *NullValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitNullValue(s)
	}
}

func (s *NullValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitNullValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) NullValue() (localctx INullValueContext) {
	localctx = NewNullValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, GraphQLParserRULE_nullValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(264)
		p.Match(GraphQLParserT__24)
	}

	return localctx
}

// IEnumValueContext is an interface to support dynamic dispatch.
type IEnumValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueContext differentiates from other interfaces.
	IsEnumValueContext()
}

type EnumValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueContext() *EnumValueContext {
	var p = new(EnumValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValue
	return p
}

func (*EnumValueContext) IsEnumValueContext() {}

func NewEnumValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueContext {
	var p = new(EnumValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValue

	return p
}

func (s *EnumValueContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValue(s)
	}
}

func (s *EnumValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValue(s)
	}
}

func (s *EnumValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitEnumValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) EnumValue() (localctx IEnumValueContext) {
	localctx = NewEnumValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, GraphQLParserRULE_enumValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(266)
		p.Name()
	}

	return localctx
}

// IListValueContext is an interface to support dynamic dispatch.
type IListValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListValueContext differentiates from other interfaces.
	IsListValueContext()
}

type ListValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListValueContext() *ListValueContext {
	var p = new(ListValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_listValue
	return p
}

func (*ListValueContext) IsListValueContext() {}

func NewListValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListValueContext {
	var p = new(ListValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_listValue

	return p
}

func (s *ListValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ListValueContext) AllValue() []IValueContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IValueContext)(nil)).Elem())
	var tst = make([]IValueContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IValueContext)
		}
	}

	return tst
}

func (s *ListValueContext) Value(i int) IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ListValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterListValue(s)
	}
}

func (s *ListValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitListValue(s)
	}
}

func (s *ListValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitListValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ListValue() (localctx IListValueContext) {
	localctx = NewListValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, GraphQLParserRULE_listValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(278)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)
			p.Match(GraphQLParserT__25)
		}
		{
			p.SetState(269)
			p.Match(GraphQLParserT__26)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)
			p.Match(GraphQLParserT__25)
		}
		p.SetState(272)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15)|(1<<GraphQLParserT__16)|(1<<GraphQLParserT__22)|(1<<GraphQLParserT__23)|(1<<GraphQLParserT__24)|(1<<GraphQLParserT__25)|(1<<GraphQLParserT__27))) != 0) || (((_la-41)&-(0x1f+1)) == 0 && ((1<<uint((_la-41)))&((1<<(GraphQLParserIntegerPart-41))|(1<<(GraphQLParserStringConst-41))|(1<<(GraphQLParserDirectiveLocation-41))|(1<<(GraphQLParserNAME-41)))) != 0) {
			{
				p.SetState(271)
				p.Value()
			}

			p.SetState(274)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(276)
			p.Match(GraphQLParserT__26)
		}

	}

	return localctx
}

// IObjectValueContext is an interface to support dynamic dispatch.
type IObjectValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectValueContext differentiates from other interfaces.
	IsObjectValueContext()
}

type ObjectValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectValueContext() *ObjectValueContext {
	var p = new(ObjectValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectValue
	return p
}

func (*ObjectValueContext) IsObjectValueContext() {}

func NewObjectValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectValueContext {
	var p = new(ObjectValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectValue

	return p
}

func (s *ObjectValueContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectValueContext) AllObjectField() []IObjectFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IObjectFieldContext)(nil)).Elem())
	var tst = make([]IObjectFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IObjectFieldContext)
		}
	}

	return tst
}

func (s *ObjectValueContext) ObjectField(i int) IObjectFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IObjectFieldContext)
}

func (s *ObjectValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectValue(s)
	}
}

func (s *ObjectValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectValue(s)
	}
}

func (s *ObjectValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitObjectValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ObjectValue() (localctx IObjectValueContext) {
	localctx = NewObjectValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, GraphQLParserRULE_objectValue)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(290)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(280)
			p.Match(GraphQLParserT__16)
		}
		{
			p.SetState(281)
			p.Match(GraphQLParserT__17)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(282)
			p.Match(GraphQLParserT__16)
		}
		p.SetState(284)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15))) != 0) || _la == GraphQLParserDirectiveLocation || _la == GraphQLParserNAME {
			{
				p.SetState(283)
				p.ObjectField()
			}

			p.SetState(286)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(288)
			p.Match(GraphQLParserT__17)
		}

	}

	return localctx
}

// IObjectFieldContext is an interface to support dynamic dispatch.
type IObjectFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectFieldContext differentiates from other interfaces.
	IsObjectFieldContext()
}

type ObjectFieldContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectFieldContext() *ObjectFieldContext {
	var p = new(ObjectFieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectField
	return p
}

func (*ObjectFieldContext) IsObjectFieldContext() {}

func NewObjectFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectFieldContext {
	var p = new(ObjectFieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectField

	return p
}

func (s *ObjectFieldContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectFieldContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectFieldContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *ObjectFieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectFieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectFieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectField(s)
	}
}

func (s *ObjectFieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectField(s)
	}
}

func (s *ObjectFieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitObjectField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ObjectField() (localctx IObjectFieldContext) {
	localctx = NewObjectFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, GraphQLParserRULE_objectField)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Name()
	}
	{
		p.SetState(293)
		p.Match(GraphQLParserT__20)
	}
	{
		p.SetState(294)
		p.Value()
	}

	return localctx
}

// IVariableContext is an interface to support dynamic dispatch.
type IVariableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableContext differentiates from other interfaces.
	IsVariableContext()
}

type VariableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableContext() *VariableContext {
	var p = new(VariableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variable
	return p
}

func (*VariableContext) IsVariableContext() {}

func NewVariableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableContext {
	var p = new(VariableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variable

	return p
}

func (s *VariableContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *VariableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariable(s)
	}
}

func (s *VariableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariable(s)
	}
}

func (s *VariableContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitVariable(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Variable() (localctx IVariableContext) {
	localctx = NewVariableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, GraphQLParserRULE_variable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(296)
		p.Match(GraphQLParserT__27)
	}
	{
		p.SetState(297)
		p.Name()
	}

	return localctx
}

// IVariableDefinitionsContext is an interface to support dynamic dispatch.
type IVariableDefinitionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDefinitionsContext differentiates from other interfaces.
	IsVariableDefinitionsContext()
}

type VariableDefinitionsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionsContext() *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinitions
	return p
}

func (*VariableDefinitionsContext) IsVariableDefinitionsContext() {}

func NewVariableDefinitionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionsContext {
	var p = new(VariableDefinitionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variableDefinitions

	return p
}

func (s *VariableDefinitionsContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionsContext) AllVariableDefinition() []IVariableDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IVariableDefinitionContext)(nil)).Elem())
	var tst = make([]IVariableDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IVariableDefinitionContext)
		}
	}

	return tst
}

func (s *VariableDefinitionsContext) VariableDefinition(i int) IVariableDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IVariableDefinitionContext)
}

func (s *VariableDefinitionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariableDefinitions(s)
	}
}

func (s *VariableDefinitionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariableDefinitions(s)
	}
}

func (s *VariableDefinitionsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitVariableDefinitions(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) VariableDefinitions() (localctx IVariableDefinitionsContext) {
	localctx = NewVariableDefinitionsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, GraphQLParserRULE_variableDefinitions)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(GraphQLParserT__18)
	}
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__27 {
		{
			p.SetState(300)
			p.VariableDefinition()
		}

		p.SetState(303)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(305)
		p.Match(GraphQLParserT__19)
	}

	return localctx
}

// IVariableDefinitionContext is an interface to support dynamic dispatch.
type IVariableDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsVariableDefinitionContext differentiates from other interfaces.
	IsVariableDefinitionContext()
}

type VariableDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVariableDefinitionContext() *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_variableDefinition
	return p
}

func (*VariableDefinitionContext) IsVariableDefinitionContext() {}

func NewVariableDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VariableDefinitionContext {
	var p = new(VariableDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_variableDefinition

	return p
}

func (s *VariableDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *VariableDefinitionContext) Variable() IVariableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVariableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVariableContext)
}

func (s *VariableDefinitionContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *VariableDefinitionContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *VariableDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VariableDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitVariableDefinition(s)
	}
}

func (s *VariableDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitVariableDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) VariableDefinition() (localctx IVariableDefinitionContext) {
	localctx = NewVariableDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, GraphQLParserRULE_variableDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(307)
		p.Variable()
	}
	{
		p.SetState(308)
		p.Match(GraphQLParserT__20)
	}
	{
		p.SetState(309)
		p.TypeSpec()
	}
	p.SetState(311)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(310)
			p.DefaultValue()
		}

	}

	return localctx
}

// IDefaultValueContext is an interface to support dynamic dispatch.
type IDefaultValueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDefaultValueContext differentiates from other interfaces.
	IsDefaultValueContext()
}

type DefaultValueContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultValueContext() *DefaultValueContext {
	var p = new(DefaultValueContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_defaultValue
	return p
}

func (*DefaultValueContext) IsDefaultValueContext() {}

func NewDefaultValueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultValueContext {
	var p = new(DefaultValueContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_defaultValue

	return p
}

func (s *DefaultValueContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultValueContext) Value() IValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IValueContext)
}

func (s *DefaultValueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultValueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultValueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDefaultValue(s)
	}
}

func (s *DefaultValueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDefaultValue(s)
	}
}

func (s *DefaultValueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDefaultValue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) DefaultValue() (localctx IDefaultValueContext) {
	localctx = NewDefaultValueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, GraphQLParserRULE_defaultValue)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(313)
		p.Value()
	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *TypeSpecContext) ListType() IListTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *TypeSpecContext) NonNullType() INonNullTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INonNullTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INonNullTypeContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSpec(s)
	}
}

func (s *TypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSpec(s)
	}
}

func (s *TypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, GraphQLParserRULE_typeSpec)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(318)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 26, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(315)
			p.NamedType()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(316)
			p.ListType()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(317)
			p.NonNullType()
		}

	}

	return localctx
}

// INamedTypeContext is an interface to support dynamic dispatch.
type INamedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNamedTypeContext differentiates from other interfaces.
	IsNamedTypeContext()
}

type NamedTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedTypeContext() *NamedTypeContext {
	var p = new(NamedTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_namedType
	return p
}

func (*NamedTypeContext) IsNamedTypeContext() {}

func NewNamedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedTypeContext {
	var p = new(NamedTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_namedType

	return p
}

func (s *NamedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedTypeContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *NamedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterNamedType(s)
	}
}

func (s *NamedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitNamedType(s)
	}
}

func (s *NamedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitNamedType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) NamedType() (localctx INamedTypeContext) {
	localctx = NewNamedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, GraphQLParserRULE_namedType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(320)
		p.Name()
	}

	return localctx
}

// IListTypeContext is an interface to support dynamic dispatch.
type IListTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsListTypeContext differentiates from other interfaces.
	IsListTypeContext()
}

type ListTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyListTypeContext() *ListTypeContext {
	var p = new(ListTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_listType
	return p
}

func (*ListTypeContext) IsListTypeContext() {}

func NewListTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListTypeContext {
	var p = new(ListTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_listType

	return p
}

func (s *ListTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ListTypeContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *ListTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterListType(s)
	}
}

func (s *ListTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitListType(s)
	}
}

func (s *ListTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitListType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ListType() (localctx IListTypeContext) {
	localctx = NewListTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, GraphQLParserRULE_listType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(322)
		p.Match(GraphQLParserT__25)
	}
	{
		p.SetState(323)
		p.TypeSpec()
	}
	{
		p.SetState(324)
		p.Match(GraphQLParserT__26)
	}

	return localctx
}

// INonNullTypeContext is an interface to support dynamic dispatch.
type INonNullTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNonNullTypeContext differentiates from other interfaces.
	IsNonNullTypeContext()
}

type NonNullTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNonNullTypeContext() *NonNullTypeContext {
	var p = new(NonNullTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_nonNullType
	return p
}

func (*NonNullTypeContext) IsNonNullTypeContext() {}

func NewNonNullTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NonNullTypeContext {
	var p = new(NonNullTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_nonNullType

	return p
}

func (s *NonNullTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *NonNullTypeContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *NonNullTypeContext) ListType() IListTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListTypeContext)
}

func (s *NonNullTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NonNullTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NonNullTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterNonNullType(s)
	}
}

func (s *NonNullTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitNonNullType(s)
	}
}

func (s *NonNullTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitNonNullType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) NonNullType() (localctx INonNullTypeContext) {
	localctx = NewNonNullTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, GraphQLParserRULE_nonNullType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(332)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case GraphQLParserT__0, GraphQLParserT__1, GraphQLParserT__2, GraphQLParserT__3, GraphQLParserT__4, GraphQLParserT__5, GraphQLParserT__6, GraphQLParserT__7, GraphQLParserT__8, GraphQLParserT__9, GraphQLParserT__10, GraphQLParserT__11, GraphQLParserT__12, GraphQLParserT__13, GraphQLParserT__14, GraphQLParserT__15, GraphQLParserDirectiveLocation, GraphQLParserNAME:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(326)
			p.NamedType()
		}
		{
			p.SetState(327)
			p.Match(GraphQLParserT__28)
		}

	case GraphQLParserT__25:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.ListType()
		}
		{
			p.SetState(330)
			p.Match(GraphQLParserT__28)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IDirectivesContext is an interface to support dynamic dispatch.
type IDirectivesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectivesContext differentiates from other interfaces.
	IsDirectivesContext()
}

type DirectivesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectivesContext() *DirectivesContext {
	var p = new(DirectivesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directives
	return p
}

func (*DirectivesContext) IsDirectivesContext() {}

func NewDirectivesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectivesContext {
	var p = new(DirectivesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directives

	return p
}

func (s *DirectivesContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectivesContext) AllDirective() []IDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveContext)(nil)).Elem())
	var tst = make([]IDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveContext)
		}
	}

	return tst
}

func (s *DirectivesContext) Directive(i int) IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *DirectivesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectivesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectivesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectives(s)
	}
}

func (s *DirectivesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectives(s)
	}
}

func (s *DirectivesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDirectives(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Directives() (localctx IDirectivesContext) {
	localctx = NewDirectivesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, GraphQLParserRULE_directives)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(335)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == GraphQLParserT__29 {
		{
			p.SetState(334)
			p.Directive()
		}

		p.SetState(337)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveContext) Arguments() IArgumentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (s *DirectiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDirective(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, GraphQLParserRULE_directive)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(339)
		p.Match(GraphQLParserT__29)
	}
	{
		p.SetState(340)
		p.Name()
	}
	p.SetState(342)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 29, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(341)
			p.Arguments()
		}

	}

	return localctx
}

// ITypeSystemDefinitionContext is an interface to support dynamic dispatch.
type ITypeSystemDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSystemDefinitionContext differentiates from other interfaces.
	IsTypeSystemDefinitionContext()
}

type TypeSystemDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSystemDefinitionContext() *TypeSystemDefinitionContext {
	var p = new(TypeSystemDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeSystemDefinition
	return p
}

func (*TypeSystemDefinitionContext) IsTypeSystemDefinitionContext() {}

func NewTypeSystemDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSystemDefinitionContext {
	var p = new(TypeSystemDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeSystemDefinition

	return p
}

func (s *TypeSystemDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSystemDefinitionContext) SchemaDefinition() ISchemaDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISchemaDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISchemaDefinitionContext)
}

func (s *TypeSystemDefinitionContext) TypeDefinition() ITypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDefinitionContext)
}

func (s *TypeSystemDefinitionContext) TypeExtension() ITypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeExtensionContext)
}

func (s *TypeSystemDefinitionContext) DirectiveDefinition() IDirectiveDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveDefinitionContext)
}

func (s *TypeSystemDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSystemDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeSystemDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeSystemDefinition(s)
	}
}

func (s *TypeSystemDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeSystemDefinition(s)
	}
}

func (s *TypeSystemDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeSystemDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeSystemDefinition() (localctx ITypeSystemDefinitionContext) {
	localctx = NewTypeSystemDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, GraphQLParserRULE_typeSystemDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(344)
			p.SchemaDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(345)
			p.TypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(346)
			p.TypeExtension()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(347)
			p.DirectiveDefinition()
		}

	}

	return localctx
}

// ISchemaDefinitionContext is an interface to support dynamic dispatch.
type ISchemaDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSchemaDefinitionContext differentiates from other interfaces.
	IsSchemaDefinitionContext()
}

type SchemaDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySchemaDefinitionContext() *SchemaDefinitionContext {
	var p = new(SchemaDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_schemaDefinition
	return p
}

func (*SchemaDefinitionContext) IsSchemaDefinitionContext() {}

func NewSchemaDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SchemaDefinitionContext {
	var p = new(SchemaDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_schemaDefinition

	return p
}

func (s *SchemaDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *SchemaDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *SchemaDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *SchemaDefinitionContext) AllOperationTypeDefinition() []IOperationTypeDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOperationTypeDefinitionContext)(nil)).Elem())
	var tst = make([]IOperationTypeDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOperationTypeDefinitionContext)
		}
	}

	return tst
}

func (s *SchemaDefinitionContext) OperationTypeDefinition(i int) IOperationTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeDefinitionContext)
}

func (s *SchemaDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SchemaDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SchemaDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterSchemaDefinition(s)
	}
}

func (s *SchemaDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitSchemaDefinition(s)
	}
}

func (s *SchemaDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitSchemaDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) SchemaDefinition() (localctx ISchemaDefinitionContext) {
	localctx = NewSchemaDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, GraphQLParserRULE_schemaDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(GraphQLParserT__5)
	}
	p.SetState(352)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15))) != 0) || _la == GraphQLParserDirectiveLocation || _la == GraphQLParserNAME {
		{
			p.SetState(351)
			p.Name()
		}

	}
	p.SetState(355)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(354)
			p.Directives()
		}

	}
	{
		p.SetState(357)
		p.Match(GraphQLParserT__16)
	}
	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14))) != 0) {
		{
			p.SetState(358)
			p.OperationTypeDefinition()
		}

		p.SetState(361)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(363)
		p.Match(GraphQLParserT__17)
	}

	return localctx
}

// IOperationTypeDefinitionContext is an interface to support dynamic dispatch.
type IOperationTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOperationTypeDefinitionContext differentiates from other interfaces.
	IsOperationTypeDefinitionContext()
}

type OperationTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOperationTypeDefinitionContext() *OperationTypeDefinitionContext {
	var p = new(OperationTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_operationTypeDefinition
	return p
}

func (*OperationTypeDefinitionContext) IsOperationTypeDefinitionContext() {}

func NewOperationTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OperationTypeDefinitionContext {
	var p = new(OperationTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_operationTypeDefinition

	return p
}

func (s *OperationTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *OperationTypeDefinitionContext) OperationType() IOperationTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOperationTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOperationTypeContext)
}

func (s *OperationTypeDefinitionContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *OperationTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OperationTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OperationTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterOperationTypeDefinition(s)
	}
}

func (s *OperationTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitOperationTypeDefinition(s)
	}
}

func (s *OperationTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitOperationTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) OperationTypeDefinition() (localctx IOperationTypeDefinitionContext) {
	localctx = NewOperationTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, GraphQLParserRULE_operationTypeDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(365)
		p.OperationType()
	}
	{
		p.SetState(366)
		p.Match(GraphQLParserT__20)
	}
	{
		p.SetState(367)
		p.NamedType()
	}

	return localctx
}

// IDescriptionContext is an interface to support dynamic dispatch.
type IDescriptionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDescriptionContext differentiates from other interfaces.
	IsDescriptionContext()
}

type DescriptionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescriptionContext() *DescriptionContext {
	var p = new(DescriptionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_description
	return p
}

func (*DescriptionContext) IsDescriptionContext() {}

func NewDescriptionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescriptionContext {
	var p = new(DescriptionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_description

	return p
}

func (s *DescriptionContext) GetParser() antlr.Parser { return s.parser }

func (s *DescriptionContext) StringValue() IStringValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringValueContext)
}

func (s *DescriptionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescriptionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescriptionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDescription(s)
	}
}

func (s *DescriptionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDescription(s)
	}
}

func (s *DescriptionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDescription(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) Description() (localctx IDescriptionContext) {
	localctx = NewDescriptionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, GraphQLParserRULE_description)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.StringValue()
	}

	return localctx
}

// ITypeDefinitionContext is an interface to support dynamic dispatch.
type ITypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeDefinitionContext differentiates from other interfaces.
	IsTypeDefinitionContext()
}

type TypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeDefinitionContext() *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeDefinition
	return p
}

func (*TypeDefinitionContext) IsTypeDefinitionContext() {}

func NewTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDefinitionContext {
	var p = new(TypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeDefinition

	return p
}

func (s *TypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDefinitionContext) ScalarTypeDefinition() IScalarTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarTypeDefinitionContext)
}

func (s *TypeDefinitionContext) ObjectTypeDefinition() IObjectTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectTypeDefinitionContext)
}

func (s *TypeDefinitionContext) InterfaceTypeDefinition() IInterfaceTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeDefinitionContext)
}

func (s *TypeDefinitionContext) UnionTypeDefinition() IUnionTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeDefinitionContext)
}

func (s *TypeDefinitionContext) EnumTypeDefinition() IEnumTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeDefinitionContext)
}

func (s *TypeDefinitionContext) InputObjectTypeDefinition() IInputObjectTypeDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputObjectTypeDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputObjectTypeDefinitionContext)
}

func (s *TypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeDefinition(s)
	}
}

func (s *TypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeDefinition() (localctx ITypeDefinitionContext) {
	localctx = NewTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, GraphQLParserRULE_typeDefinition)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(377)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(371)
			p.ScalarTypeDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(372)
			p.ObjectTypeDefinition()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(373)
			p.InterfaceTypeDefinition()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(374)
			p.UnionTypeDefinition()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(375)
			p.EnumTypeDefinition()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(376)
			p.InputObjectTypeDefinition()
		}

	}

	return localctx
}

// ITypeExtensionContext is an interface to support dynamic dispatch.
type ITypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeExtensionContext differentiates from other interfaces.
	IsTypeExtensionContext()
}

type TypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeExtensionContext() *TypeExtensionContext {
	var p = new(TypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_typeExtension
	return p
}

func (*TypeExtensionContext) IsTypeExtensionContext() {}

func NewTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeExtensionContext {
	var p = new(TypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_typeExtension

	return p
}

func (s *TypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeExtensionContext) ScalarTypeExtension() IScalarTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IScalarTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IScalarTypeExtensionContext)
}

func (s *TypeExtensionContext) ObjectTypeExtension() IObjectTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IObjectTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IObjectTypeExtensionContext)
}

func (s *TypeExtensionContext) InterfaceTypeExtension() IInterfaceTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInterfaceTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInterfaceTypeExtensionContext)
}

func (s *TypeExtensionContext) UnionTypeExtension() IUnionTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionTypeExtensionContext)
}

func (s *TypeExtensionContext) EnumTypeExtension() IEnumTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumTypeExtensionContext)
}

func (s *TypeExtensionContext) InputObjectTypeExtension() IInputObjectTypeExtensionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputObjectTypeExtensionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputObjectTypeExtensionContext)
}

func (s *TypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterTypeExtension(s)
	}
}

func (s *TypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitTypeExtension(s)
	}
}

func (s *TypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) TypeExtension() (localctx ITypeExtensionContext) {
	localctx = NewTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, GraphQLParserRULE_typeExtension)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(379)
			p.ScalarTypeExtension()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(380)
			p.ObjectTypeExtension()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(381)
			p.InterfaceTypeExtension()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(382)
			p.UnionTypeExtension()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(383)
			p.EnumTypeExtension()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(384)
			p.InputObjectTypeExtension()
		}

	}

	return localctx
}

// IScalarTypeDefinitionContext is an interface to support dynamic dispatch.
type IScalarTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarTypeDefinitionContext differentiates from other interfaces.
	IsScalarTypeDefinitionContext()
}

type ScalarTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeDefinitionContext() *ScalarTypeDefinitionContext {
	var p = new(ScalarTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_scalarTypeDefinition
	return p
}

func (*ScalarTypeDefinitionContext) IsScalarTypeDefinitionContext() {}

func NewScalarTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeDefinitionContext {
	var p = new(ScalarTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_scalarTypeDefinition

	return p
}

func (s *ScalarTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ScalarTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ScalarTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ScalarTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterScalarTypeDefinition(s)
	}
}

func (s *ScalarTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitScalarTypeDefinition(s)
	}
}

func (s *ScalarTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitScalarTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ScalarTypeDefinition() (localctx IScalarTypeDefinitionContext) {
	localctx = NewScalarTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, GraphQLParserRULE_scalarTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(388)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringConst {
		{
			p.SetState(387)
			p.Description()
		}

	}
	{
		p.SetState(390)
		p.Match(GraphQLParserT__10)
	}
	{
		p.SetState(391)
		p.Name()
	}
	p.SetState(393)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(392)
			p.Directives()
		}

	}

	return localctx
}

// IScalarTypeExtensionContext is an interface to support dynamic dispatch.
type IScalarTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsScalarTypeExtensionContext differentiates from other interfaces.
	IsScalarTypeExtensionContext()
}

type ScalarTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyScalarTypeExtensionContext() *ScalarTypeExtensionContext {
	var p = new(ScalarTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_scalarTypeExtension
	return p
}

func (*ScalarTypeExtensionContext) IsScalarTypeExtensionContext() {}

func NewScalarTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ScalarTypeExtensionContext {
	var p = new(ScalarTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_scalarTypeExtension

	return p
}

func (s *ScalarTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ScalarTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ScalarTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ScalarTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ScalarTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ScalarTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterScalarTypeExtension(s)
	}
}

func (s *ScalarTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitScalarTypeExtension(s)
	}
}

func (s *ScalarTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitScalarTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ScalarTypeExtension() (localctx IScalarTypeExtensionContext) {
	localctx = NewScalarTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, GraphQLParserRULE_scalarTypeExtension)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(395)
		p.Match(GraphQLParserT__3)
	}
	{
		p.SetState(396)
		p.Match(GraphQLParserT__10)
	}
	{
		p.SetState(397)
		p.Name()
	}
	{
		p.SetState(398)
		p.Directives()
	}

	return localctx
}

// IObjectTypeDefinitionContext is an interface to support dynamic dispatch.
type IObjectTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectTypeDefinitionContext differentiates from other interfaces.
	IsObjectTypeDefinitionContext()
}

type ObjectTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeDefinitionContext() *ObjectTypeDefinitionContext {
	var p = new(ObjectTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectTypeDefinition
	return p
}

func (*ObjectTypeDefinitionContext) IsObjectTypeDefinitionContext() {}

func NewObjectTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeDefinitionContext {
	var p = new(ObjectTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectTypeDefinition

	return p
}

func (s *ObjectTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *ObjectTypeDefinitionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ObjectTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ObjectTypeDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *ObjectTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectTypeDefinition(s)
	}
}

func (s *ObjectTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectTypeDefinition(s)
	}
}

func (s *ObjectTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitObjectTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ObjectTypeDefinition() (localctx IObjectTypeDefinitionContext) {
	localctx = NewObjectTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, GraphQLParserRULE_objectTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringConst {
		{
			p.SetState(400)
			p.Description()
		}

	}
	{
		p.SetState(403)
		p.Match(GraphQLParserT__0)
	}
	{
		p.SetState(404)
		p.Name()
	}
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__4 {
		{
			p.SetState(405)
			p.implementsInterfaces(0)
		}

	}
	p.SetState(409)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(408)
			p.Directives()
		}

	}
	p.SetState(412)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(411)
			p.FieldsDefinition()
		}

	}

	return localctx
}

// IObjectTypeExtensionContext is an interface to support dynamic dispatch.
type IObjectTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsObjectTypeExtensionContext differentiates from other interfaces.
	IsObjectTypeExtensionContext()
}

type ObjectTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectTypeExtensionContext() *ObjectTypeExtensionContext {
	var p = new(ObjectTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_objectTypeExtension
	return p
}

func (*ObjectTypeExtensionContext) IsObjectTypeExtensionContext() {}

func NewObjectTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectTypeExtensionContext {
	var p = new(ObjectTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_objectTypeExtension

	return p
}

func (s *ObjectTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectTypeExtensionContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *ObjectTypeExtensionContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *ObjectTypeExtensionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *ObjectTypeExtensionContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ObjectTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *ObjectTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterObjectTypeExtension(s)
	}
}

func (s *ObjectTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitObjectTypeExtension(s)
	}
}

func (s *ObjectTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitObjectTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ObjectTypeExtension() (localctx IObjectTypeExtensionContext) {
	localctx = NewObjectTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, GraphQLParserRULE_objectTypeExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(449)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(414)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(415)
			p.Match(GraphQLParserT__0)
		}
		{
			p.SetState(416)
			p.Name()
		}
		p.SetState(418)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__4 {
			{
				p.SetState(417)
				p.implementsInterfaces(0)
			}

		}
		p.SetState(421)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__29 {
			{
				p.SetState(420)
				p.Directives()
			}

		}
		{
			p.SetState(423)
			p.FieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(425)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(426)
			p.Match(GraphQLParserT__0)
		}
		{
			p.SetState(427)
			p.Name()
		}
		p.SetState(429)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__4 {
			{
				p.SetState(428)
				p.implementsInterfaces(0)
			}

		}
		{
			p.SetState(431)
			p.Directives()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(433)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(434)
			p.Match(GraphQLParserT__0)
		}
		{
			p.SetState(435)
			p.Name()
		}
		{
			p.SetState(436)
			p.implementsInterfaces(0)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(438)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(439)
			p.Match(GraphQLParserT__0)
		}
		{
			p.SetState(440)
			p.Name()
		}
		{
			p.SetState(441)
			p.Match(GraphQLParserT__15)
		}
		{
			p.SetState(442)
			p.Name()
		}
		p.SetState(444)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__29 {
			{
				p.SetState(443)
				p.Directives()
			}

		}
		p.SetState(447)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(446)
				p.FieldsDefinition()
			}

		}

	}

	return localctx
}

// IImplementsInterfacesContext is an interface to support dynamic dispatch.
type IImplementsInterfacesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsImplementsInterfacesContext differentiates from other interfaces.
	IsImplementsInterfacesContext()
}

type ImplementsInterfacesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyImplementsInterfacesContext() *ImplementsInterfacesContext {
	var p = new(ImplementsInterfacesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_implementsInterfaces
	return p
}

func (*ImplementsInterfacesContext) IsImplementsInterfacesContext() {}

func NewImplementsInterfacesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImplementsInterfacesContext {
	var p = new(ImplementsInterfacesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_implementsInterfaces

	return p
}

func (s *ImplementsInterfacesContext) GetParser() antlr.Parser { return s.parser }

func (s *ImplementsInterfacesContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *ImplementsInterfacesContext) ImplementsInterfaces() IImplementsInterfacesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImplementsInterfacesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImplementsInterfacesContext)
}

func (s *ImplementsInterfacesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImplementsInterfacesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImplementsInterfacesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterImplementsInterfaces(s)
	}
}

func (s *ImplementsInterfacesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitImplementsInterfaces(s)
	}
}

func (s *ImplementsInterfacesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitImplementsInterfaces(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ImplementsInterfaces() (localctx IImplementsInterfacesContext) {
	return p.implementsInterfaces(0)
}

func (p *GraphQLParser) implementsInterfaces(_p int) (localctx IImplementsInterfacesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewImplementsInterfacesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IImplementsInterfacesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 96
	p.EnterRecursionRule(localctx, 96, GraphQLParserRULE_implementsInterfaces, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(452)
		p.Match(GraphQLParserT__4)
	}
	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__30 {
		{
			p.SetState(453)
			p.Match(GraphQLParserT__30)
		}

	}
	{
		p.SetState(456)
		p.NamedType()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(465)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewImplementsInterfacesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GraphQLParserRULE_implementsInterfaces)
			p.SetState(458)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			p.SetState(460)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)

			if _la == GraphQLParserT__30 {
				{
					p.SetState(459)
					p.Match(GraphQLParserT__30)
				}

			}
			{
				p.SetState(462)
				p.NamedType()
			}

		}
		p.SetState(467)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 50, p.GetParserRuleContext())
	}

	return localctx
}

// IFieldsDefinitionContext is an interface to support dynamic dispatch.
type IFieldsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldsDefinitionContext differentiates from other interfaces.
	IsFieldsDefinitionContext()
}

type FieldsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldsDefinitionContext() *FieldsDefinitionContext {
	var p = new(FieldsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldsDefinition
	return p
}

func (*FieldsDefinitionContext) IsFieldsDefinitionContext() {}

func NewFieldsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldsDefinitionContext {
	var p = new(FieldsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fieldsDefinition

	return p
}

func (s *FieldsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldsDefinitionContext) AllFieldDefinition() []IFieldDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldDefinitionContext)(nil)).Elem())
	var tst = make([]IFieldDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldDefinitionContext)
		}
	}

	return tst
}

func (s *FieldsDefinitionContext) FieldDefinition(i int) IFieldDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldDefinitionContext)
}

func (s *FieldsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFieldsDefinition(s)
	}
}

func (s *FieldsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFieldsDefinition(s)
	}
}

func (s *FieldsDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFieldsDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FieldsDefinition() (localctx IFieldsDefinitionContext) {
	localctx = NewFieldsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, GraphQLParserRULE_fieldsDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(468)
		p.Match(GraphQLParserT__16)
	}
	p.SetState(470)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15))) != 0) || (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(GraphQLParserStringConst-45))|(1<<(GraphQLParserDirectiveLocation-45))|(1<<(GraphQLParserNAME-45)))) != 0) {
		{
			p.SetState(469)
			p.FieldDefinition()
		}

		p.SetState(472)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(474)
		p.Match(GraphQLParserT__17)
	}

	return localctx
}

// IFieldDefinitionContext is an interface to support dynamic dispatch.
type IFieldDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFieldDefinitionContext differentiates from other interfaces.
	IsFieldDefinitionContext()
}

type FieldDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFieldDefinitionContext() *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_fieldDefinition
	return p
}

func (*FieldDefinitionContext) IsFieldDefinitionContext() {}

func NewFieldDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldDefinitionContext {
	var p = new(FieldDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_fieldDefinition

	return p
}

func (s *FieldDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *FieldDefinitionContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *FieldDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *FieldDefinitionContext) ArgumentsDefinition() IArgumentsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsDefinitionContext)
}

func (s *FieldDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *FieldDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterFieldDefinition(s)
	}
}

func (s *FieldDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitFieldDefinition(s)
	}
}

func (s *FieldDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitFieldDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) FieldDefinition() (localctx IFieldDefinitionContext) {
	localctx = NewFieldDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, GraphQLParserRULE_fieldDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(477)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringConst {
		{
			p.SetState(476)
			p.Description()
		}

	}
	{
		p.SetState(479)
		p.Name()
	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__18 {
		{
			p.SetState(480)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(483)
		p.Match(GraphQLParserT__20)
	}
	{
		p.SetState(484)
		p.TypeSpec()
	}
	p.SetState(486)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(485)
			p.Directives()
		}

	}

	return localctx
}

// IArgumentsDefinitionContext is an interface to support dynamic dispatch.
type IArgumentsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsArgumentsDefinitionContext differentiates from other interfaces.
	IsArgumentsDefinitionContext()
}

type ArgumentsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgumentsDefinitionContext() *ArgumentsDefinitionContext {
	var p = new(ArgumentsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_argumentsDefinition
	return p
}

func (*ArgumentsDefinitionContext) IsArgumentsDefinitionContext() {}

func NewArgumentsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgumentsDefinitionContext {
	var p = new(ArgumentsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_argumentsDefinition

	return p
}

func (s *ArgumentsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgumentsDefinitionContext) AllInputValueDefinition() []IInputValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem())
	var tst = make([]IInputValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInputValueDefinitionContext)
		}
	}

	return tst
}

func (s *ArgumentsDefinitionContext) InputValueDefinition(i int) IInputValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInputValueDefinitionContext)
}

func (s *ArgumentsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgumentsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgumentsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterArgumentsDefinition(s)
	}
}

func (s *ArgumentsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitArgumentsDefinition(s)
	}
}

func (s *ArgumentsDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitArgumentsDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) ArgumentsDefinition() (localctx IArgumentsDefinitionContext) {
	localctx = NewArgumentsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, GraphQLParserRULE_argumentsDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(488)
		p.Match(GraphQLParserT__18)
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15))) != 0) || (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(GraphQLParserStringConst-45))|(1<<(GraphQLParserDirectiveLocation-45))|(1<<(GraphQLParserNAME-45)))) != 0) {
		{
			p.SetState(489)
			p.InputValueDefinition()
		}

		p.SetState(494)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(495)
		p.Match(GraphQLParserT__19)
	}

	return localctx
}

// IInputValueDefinitionContext is an interface to support dynamic dispatch.
type IInputValueDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputValueDefinitionContext differentiates from other interfaces.
	IsInputValueDefinitionContext()
}

type InputValueDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputValueDefinitionContext() *InputValueDefinitionContext {
	var p = new(InputValueDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputValueDefinition
	return p
}

func (*InputValueDefinitionContext) IsInputValueDefinitionContext() {}

func NewInputValueDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputValueDefinitionContext {
	var p = new(InputValueDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputValueDefinition

	return p
}

func (s *InputValueDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputValueDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputValueDefinitionContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *InputValueDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InputValueDefinitionContext) DefaultValue() IDefaultValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultValueContext)
}

func (s *InputValueDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputValueDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputValueDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputValueDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputValueDefinition(s)
	}
}

func (s *InputValueDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputValueDefinition(s)
	}
}

func (s *InputValueDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInputValueDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InputValueDefinition() (localctx IInputValueDefinitionContext) {
	localctx = NewInputValueDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, GraphQLParserRULE_inputValueDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(498)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringConst {
		{
			p.SetState(497)
			p.Description()
		}

	}
	{
		p.SetState(500)
		p.Name()
	}
	{
		p.SetState(501)
		p.Match(GraphQLParserT__20)
	}
	{
		p.SetState(502)
		p.TypeSpec()
	}
	p.SetState(507)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) == 1 {
		p.SetState(504)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__31 {
			{
				p.SetState(503)
				p.Match(GraphQLParserT__31)
			}

		}
		{
			p.SetState(506)
			p.DefaultValue()
		}

	}
	p.SetState(510)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(509)
			p.Directives()
		}

	}

	return localctx
}

// IInterfaceTypeDefinitionContext is an interface to support dynamic dispatch.
type IInterfaceTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceTypeDefinitionContext differentiates from other interfaces.
	IsInterfaceTypeDefinitionContext()
}

type InterfaceTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeDefinitionContext() *InterfaceTypeDefinitionContext {
	var p = new(InterfaceTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_interfaceTypeDefinition
	return p
}

func (*InterfaceTypeDefinitionContext) IsInterfaceTypeDefinitionContext() {}

func NewInterfaceTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeDefinitionContext {
	var p = new(InterfaceTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_interfaceTypeDefinition

	return p
}

func (s *InterfaceTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InterfaceTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InterfaceTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InterfaceTypeDefinitionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *InterfaceTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInterfaceTypeDefinition(s)
	}
}

func (s *InterfaceTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInterfaceTypeDefinition(s)
	}
}

func (s *InterfaceTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInterfaceTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InterfaceTypeDefinition() (localctx IInterfaceTypeDefinitionContext) {
	localctx = NewInterfaceTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, GraphQLParserRULE_interfaceTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(513)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringConst {
		{
			p.SetState(512)
			p.Description()
		}

	}
	{
		p.SetState(515)
		p.Match(GraphQLParserT__8)
	}
	{
		p.SetState(516)
		p.Name()
	}
	p.SetState(518)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(517)
			p.Directives()
		}

	}
	p.SetState(521)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 62, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(520)
			p.FieldsDefinition()
		}

	}

	return localctx
}

// IInterfaceTypeExtensionContext is an interface to support dynamic dispatch.
type IInterfaceTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInterfaceTypeExtensionContext differentiates from other interfaces.
	IsInterfaceTypeExtensionContext()
}

type InterfaceTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInterfaceTypeExtensionContext() *InterfaceTypeExtensionContext {
	var p = new(InterfaceTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_interfaceTypeExtension
	return p
}

func (*InterfaceTypeExtensionContext) IsInterfaceTypeExtensionContext() {}

func NewInterfaceTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InterfaceTypeExtensionContext {
	var p = new(InterfaceTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_interfaceTypeExtension

	return p
}

func (s *InterfaceTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *InterfaceTypeExtensionContext) AllName() []INameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INameContext)(nil)).Elem())
	var tst = make([]INameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INameContext)
		}
	}

	return tst
}

func (s *InterfaceTypeExtensionContext) Name(i int) INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InterfaceTypeExtensionContext) FieldsDefinition() IFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldsDefinitionContext)
}

func (s *InterfaceTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InterfaceTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InterfaceTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInterfaceTypeExtension(s)
	}
}

func (s *InterfaceTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInterfaceTypeExtension(s)
	}
}

func (s *InterfaceTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInterfaceTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InterfaceTypeExtension() (localctx IInterfaceTypeExtensionContext) {
	localctx = NewInterfaceTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, GraphQLParserRULE_interfaceTypeExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(547)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 66, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(523)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(524)
			p.Match(GraphQLParserT__8)
		}
		{
			p.SetState(525)
			p.Name()
		}
		p.SetState(527)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__29 {
			{
				p.SetState(526)
				p.Directives()
			}

		}
		{
			p.SetState(529)
			p.FieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(531)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(532)
			p.Match(GraphQLParserT__8)
		}
		{
			p.SetState(533)
			p.Name()
		}
		{
			p.SetState(534)
			p.Directives()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(536)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(537)
			p.Match(GraphQLParserT__8)
		}
		{
			p.SetState(538)
			p.Name()
		}
		{
			p.SetState(539)
			p.Match(GraphQLParserT__15)
		}
		{
			p.SetState(540)
			p.Name()
		}
		p.SetState(542)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__29 {
			{
				p.SetState(541)
				p.Directives()
			}

		}
		p.SetState(545)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 65, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(544)
				p.FieldsDefinition()
			}

		}

	}

	return localctx
}

// IUnionTypeDefinitionContext is an interface to support dynamic dispatch.
type IUnionTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeDefinitionContext differentiates from other interfaces.
	IsUnionTypeDefinitionContext()
}

type UnionTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeDefinitionContext() *UnionTypeDefinitionContext {
	var p = new(UnionTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionTypeDefinition
	return p
}

func (*UnionTypeDefinitionContext) IsUnionTypeDefinitionContext() {}

func NewUnionTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeDefinitionContext {
	var p = new(UnionTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionTypeDefinition

	return p
}

func (s *UnionTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnionTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *UnionTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *UnionTypeDefinitionContext) UnionMemberTypes() IUnionMemberTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMemberTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMemberTypesContext)
}

func (s *UnionTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionTypeDefinition(s)
	}
}

func (s *UnionTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionTypeDefinition(s)
	}
}

func (s *UnionTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitUnionTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) UnionTypeDefinition() (localctx IUnionTypeDefinitionContext) {
	localctx = NewUnionTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, GraphQLParserRULE_unionTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringConst {
		{
			p.SetState(549)
			p.Description()
		}

	}
	{
		p.SetState(552)
		p.Match(GraphQLParserT__7)
	}
	{
		p.SetState(553)
		p.Name()
	}
	p.SetState(555)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(554)
			p.Directives()
		}

	}
	p.SetState(558)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__31 {
		{
			p.SetState(557)
			p.unionMemberTypes(0)
		}

	}

	return localctx
}

// IUnionMemberTypesContext is an interface to support dynamic dispatch.
type IUnionMemberTypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionMemberTypesContext differentiates from other interfaces.
	IsUnionMemberTypesContext()
}

type UnionMemberTypesContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionMemberTypesContext() *UnionMemberTypesContext {
	var p = new(UnionMemberTypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionMemberTypes
	return p
}

func (*UnionMemberTypesContext) IsUnionMemberTypesContext() {}

func NewUnionMemberTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionMemberTypesContext {
	var p = new(UnionMemberTypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionMemberTypes

	return p
}

func (s *UnionMemberTypesContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionMemberTypesContext) NamedType() INamedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INamedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INamedTypeContext)
}

func (s *UnionMemberTypesContext) UnionMemberTypes() IUnionMemberTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMemberTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMemberTypesContext)
}

func (s *UnionMemberTypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionMemberTypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionMemberTypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionMemberTypes(s)
	}
}

func (s *UnionMemberTypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionMemberTypes(s)
	}
}

func (s *UnionMemberTypesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitUnionMemberTypes(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) UnionMemberTypes() (localctx IUnionMemberTypesContext) {
	return p.unionMemberTypes(0)
}

func (p *GraphQLParser) unionMemberTypes(_p int) (localctx IUnionMemberTypesContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewUnionMemberTypesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IUnionMemberTypesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 112
	p.EnterRecursionRule(localctx, 112, GraphQLParserRULE_unionMemberTypes, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(561)
		p.Match(GraphQLParserT__31)
	}
	p.SetState(563)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__32 {
		{
			p.SetState(562)
			p.Match(GraphQLParserT__32)
		}

	}
	{
		p.SetState(565)
		p.NamedType()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(572)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewUnionMemberTypesContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GraphQLParserRULE_unionMemberTypes)
			p.SetState(567)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(568)
				p.Match(GraphQLParserT__32)
			}
			{
				p.SetState(569)
				p.NamedType()
			}

		}
		p.SetState(574)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 71, p.GetParserRuleContext())
	}

	return localctx
}

// IUnionTypeExtensionContext is an interface to support dynamic dispatch.
type IUnionTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnionTypeExtensionContext differentiates from other interfaces.
	IsUnionTypeExtensionContext()
}

type UnionTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnionTypeExtensionContext() *UnionTypeExtensionContext {
	var p = new(UnionTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_unionTypeExtension
	return p
}

func (*UnionTypeExtensionContext) IsUnionTypeExtensionContext() {}

func NewUnionTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnionTypeExtensionContext {
	var p = new(UnionTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_unionTypeExtension

	return p
}

func (s *UnionTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnionTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *UnionTypeExtensionContext) UnionMemberTypes() IUnionMemberTypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnionMemberTypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnionMemberTypesContext)
}

func (s *UnionTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *UnionTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnionTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnionTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterUnionTypeExtension(s)
	}
}

func (s *UnionTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitUnionTypeExtension(s)
	}
}

func (s *UnionTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitUnionTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) UnionTypeExtension() (localctx IUnionTypeExtensionContext) {
	localctx = NewUnionTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, GraphQLParserRULE_unionTypeExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(588)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 73, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(575)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(576)
			p.Match(GraphQLParserT__7)
		}
		{
			p.SetState(577)
			p.Name()
		}
		p.SetState(579)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__29 {
			{
				p.SetState(578)
				p.Directives()
			}

		}
		{
			p.SetState(581)
			p.unionMemberTypes(0)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(583)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(584)
			p.Match(GraphQLParserT__7)
		}
		{
			p.SetState(585)
			p.Name()
		}
		{
			p.SetState(586)
			p.Directives()
		}

	}

	return localctx
}

// IEnumTypeDefinitionContext is an interface to support dynamic dispatch.
type IEnumTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeDefinitionContext differentiates from other interfaces.
	IsEnumTypeDefinitionContext()
}

type EnumTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeDefinitionContext() *EnumTypeDefinitionContext {
	var p = new(EnumTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumTypeDefinition
	return p
}

func (*EnumTypeDefinitionContext) IsEnumTypeDefinitionContext() {}

func NewEnumTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeDefinitionContext {
	var p = new(EnumTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumTypeDefinition

	return p
}

func (s *EnumTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EnumTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumTypeDefinitionContext) EnumValuesDefinition() IEnumValuesDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValuesDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValuesDefinitionContext)
}

func (s *EnumTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumTypeDefinition(s)
	}
}

func (s *EnumTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumTypeDefinition(s)
	}
}

func (s *EnumTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitEnumTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) EnumTypeDefinition() (localctx IEnumTypeDefinitionContext) {
	localctx = NewEnumTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, GraphQLParserRULE_enumTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(591)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringConst {
		{
			p.SetState(590)
			p.Description()
		}

	}
	{
		p.SetState(593)
		p.Match(GraphQLParserT__6)
	}
	{
		p.SetState(594)
		p.Name()
	}
	p.SetState(596)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(595)
			p.Directives()
		}

	}
	p.SetState(599)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 76, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(598)
			p.EnumValuesDefinition()
		}

	}

	return localctx
}

// IEnumValuesDefinitionContext is an interface to support dynamic dispatch.
type IEnumValuesDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValuesDefinitionContext differentiates from other interfaces.
	IsEnumValuesDefinitionContext()
}

type EnumValuesDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValuesDefinitionContext() *EnumValuesDefinitionContext {
	var p = new(EnumValuesDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValuesDefinition
	return p
}

func (*EnumValuesDefinitionContext) IsEnumValuesDefinitionContext() {}

func NewEnumValuesDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValuesDefinitionContext {
	var p = new(EnumValuesDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValuesDefinition

	return p
}

func (s *EnumValuesDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValuesDefinitionContext) AllEnumValueDefinition() []IEnumValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IEnumValueDefinitionContext)(nil)).Elem())
	var tst = make([]IEnumValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IEnumValueDefinitionContext)
		}
	}

	return tst
}

func (s *EnumValuesDefinitionContext) EnumValueDefinition(i int) IEnumValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IEnumValueDefinitionContext)
}

func (s *EnumValuesDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValuesDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValuesDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValuesDefinition(s)
	}
}

func (s *EnumValuesDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValuesDefinition(s)
	}
}

func (s *EnumValuesDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitEnumValuesDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) EnumValuesDefinition() (localctx IEnumValuesDefinitionContext) {
	localctx = NewEnumValuesDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, GraphQLParserRULE_enumValuesDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(601)
		p.Match(GraphQLParserT__16)
	}
	p.SetState(603)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15))) != 0) || (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(GraphQLParserStringConst-45))|(1<<(GraphQLParserDirectiveLocation-45))|(1<<(GraphQLParserNAME-45)))) != 0) {
		{
			p.SetState(602)
			p.EnumValueDefinition()
		}

		p.SetState(605)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(607)
		p.Match(GraphQLParserT__17)
	}

	return localctx
}

// IEnumValueDefinitionContext is an interface to support dynamic dispatch.
type IEnumValueDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumValueDefinitionContext differentiates from other interfaces.
	IsEnumValueDefinitionContext()
}

type EnumValueDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumValueDefinitionContext() *EnumValueDefinitionContext {
	var p = new(EnumValueDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumValueDefinition
	return p
}

func (*EnumValueDefinitionContext) IsEnumValueDefinitionContext() {}

func NewEnumValueDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumValueDefinitionContext {
	var p = new(EnumValueDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumValueDefinition

	return p
}

func (s *EnumValueDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumValueDefinitionContext) EnumValue() IEnumValueContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValueContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValueContext)
}

func (s *EnumValueDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *EnumValueDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumValueDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumValueDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumValueDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumValueDefinition(s)
	}
}

func (s *EnumValueDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumValueDefinition(s)
	}
}

func (s *EnumValueDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitEnumValueDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) EnumValueDefinition() (localctx IEnumValueDefinitionContext) {
	localctx = NewEnumValueDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, GraphQLParserRULE_enumValueDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(610)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringConst {
		{
			p.SetState(609)
			p.Description()
		}

	}
	{
		p.SetState(612)
		p.EnumValue()
	}
	p.SetState(614)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(613)
			p.Directives()
		}

	}

	return localctx
}

// IEnumTypeExtensionContext is an interface to support dynamic dispatch.
type IEnumTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsEnumTypeExtensionContext differentiates from other interfaces.
	IsEnumTypeExtensionContext()
}

type EnumTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEnumTypeExtensionContext() *EnumTypeExtensionContext {
	var p = new(EnumTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_enumTypeExtension
	return p
}

func (*EnumTypeExtensionContext) IsEnumTypeExtensionContext() {}

func NewEnumTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EnumTypeExtensionContext {
	var p = new(EnumTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_enumTypeExtension

	return p
}

func (s *EnumTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *EnumTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *EnumTypeExtensionContext) EnumValuesDefinition() IEnumValuesDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IEnumValuesDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IEnumValuesDefinitionContext)
}

func (s *EnumTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *EnumTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EnumTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EnumTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterEnumTypeExtension(s)
	}
}

func (s *EnumTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitEnumTypeExtension(s)
	}
}

func (s *EnumTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitEnumTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) EnumTypeExtension() (localctx IEnumTypeExtensionContext) {
	localctx = NewEnumTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, GraphQLParserRULE_enumTypeExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(629)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 81, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(616)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(617)
			p.Match(GraphQLParserT__6)
		}
		{
			p.SetState(618)
			p.Name()
		}
		p.SetState(620)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__29 {
			{
				p.SetState(619)
				p.Directives()
			}

		}
		{
			p.SetState(622)
			p.EnumValuesDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(624)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(625)
			p.Match(GraphQLParserT__6)
		}
		{
			p.SetState(626)
			p.Name()
		}
		{
			p.SetState(627)
			p.Directives()
		}

	}

	return localctx
}

// IInputObjectTypeDefinitionContext is an interface to support dynamic dispatch.
type IInputObjectTypeDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputObjectTypeDefinitionContext differentiates from other interfaces.
	IsInputObjectTypeDefinitionContext()
}

type InputObjectTypeDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectTypeDefinitionContext() *InputObjectTypeDefinitionContext {
	var p = new(InputObjectTypeDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeDefinition
	return p
}

func (*InputObjectTypeDefinitionContext) IsInputObjectTypeDefinitionContext() {}

func NewInputObjectTypeDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectTypeDefinitionContext {
	var p = new(InputObjectTypeDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeDefinition

	return p
}

func (s *InputObjectTypeDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectTypeDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputObjectTypeDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *InputObjectTypeDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputObjectTypeDefinitionContext) InputFieldsDefinition() IInputFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputFieldsDefinitionContext)
}

func (s *InputObjectTypeDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectTypeDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectTypeDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputObjectTypeDefinition(s)
	}
}

func (s *InputObjectTypeDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputObjectTypeDefinition(s)
	}
}

func (s *InputObjectTypeDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInputObjectTypeDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InputObjectTypeDefinition() (localctx IInputObjectTypeDefinitionContext) {
	localctx = NewInputObjectTypeDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, GraphQLParserRULE_inputObjectTypeDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(632)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringConst {
		{
			p.SetState(631)
			p.Description()
		}

	}
	{
		p.SetState(634)
		p.Match(GraphQLParserT__9)
	}
	{
		p.SetState(635)
		p.Name()
	}
	p.SetState(637)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(636)
			p.Directives()
		}

	}
	p.SetState(640)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 84, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(639)
			p.InputFieldsDefinition()
		}

	}

	return localctx
}

// IInputFieldsDefinitionContext is an interface to support dynamic dispatch.
type IInputFieldsDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputFieldsDefinitionContext differentiates from other interfaces.
	IsInputFieldsDefinitionContext()
}

type InputFieldsDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputFieldsDefinitionContext() *InputFieldsDefinitionContext {
	var p = new(InputFieldsDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputFieldsDefinition
	return p
}

func (*InputFieldsDefinitionContext) IsInputFieldsDefinitionContext() {}

func NewInputFieldsDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputFieldsDefinitionContext {
	var p = new(InputFieldsDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputFieldsDefinition

	return p
}

func (s *InputFieldsDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputFieldsDefinitionContext) AllInputValueDefinition() []IInputValueDefinitionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem())
	var tst = make([]IInputValueDefinitionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInputValueDefinitionContext)
		}
	}

	return tst
}

func (s *InputFieldsDefinitionContext) InputValueDefinition(i int) IInputValueDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputValueDefinitionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInputValueDefinitionContext)
}

func (s *InputFieldsDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputFieldsDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputFieldsDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputFieldsDefinition(s)
	}
}

func (s *InputFieldsDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputFieldsDefinition(s)
	}
}

func (s *InputFieldsDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInputFieldsDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InputFieldsDefinition() (localctx IInputFieldsDefinitionContext) {
	localctx = NewInputFieldsDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, GraphQLParserRULE_inputFieldsDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(642)
		p.Match(GraphQLParserT__16)
	}
	p.SetState(644)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<GraphQLParserT__0)|(1<<GraphQLParserT__1)|(1<<GraphQLParserT__2)|(1<<GraphQLParserT__3)|(1<<GraphQLParserT__4)|(1<<GraphQLParserT__5)|(1<<GraphQLParserT__6)|(1<<GraphQLParserT__7)|(1<<GraphQLParserT__8)|(1<<GraphQLParserT__9)|(1<<GraphQLParserT__10)|(1<<GraphQLParserT__11)|(1<<GraphQLParserT__12)|(1<<GraphQLParserT__13)|(1<<GraphQLParserT__14)|(1<<GraphQLParserT__15))) != 0) || (((_la-45)&-(0x1f+1)) == 0 && ((1<<uint((_la-45)))&((1<<(GraphQLParserStringConst-45))|(1<<(GraphQLParserDirectiveLocation-45))|(1<<(GraphQLParserNAME-45)))) != 0) {
		{
			p.SetState(643)
			p.InputValueDefinition()
		}

		p.SetState(646)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(648)
		p.Match(GraphQLParserT__17)
	}

	return localctx
}

// IInputObjectTypeExtensionContext is an interface to support dynamic dispatch.
type IInputObjectTypeExtensionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInputObjectTypeExtensionContext differentiates from other interfaces.
	IsInputObjectTypeExtensionContext()
}

type InputObjectTypeExtensionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInputObjectTypeExtensionContext() *InputObjectTypeExtensionContext {
	var p = new(InputObjectTypeExtensionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeExtension
	return p
}

func (*InputObjectTypeExtensionContext) IsInputObjectTypeExtensionContext() {}

func NewInputObjectTypeExtensionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InputObjectTypeExtensionContext {
	var p = new(InputObjectTypeExtensionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_inputObjectTypeExtension

	return p
}

func (s *InputObjectTypeExtensionContext) GetParser() antlr.Parser { return s.parser }

func (s *InputObjectTypeExtensionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *InputObjectTypeExtensionContext) InputFieldsDefinition() IInputFieldsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInputFieldsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInputFieldsDefinitionContext)
}

func (s *InputObjectTypeExtensionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *InputObjectTypeExtensionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InputObjectTypeExtensionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InputObjectTypeExtensionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterInputObjectTypeExtension(s)
	}
}

func (s *InputObjectTypeExtensionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitInputObjectTypeExtension(s)
	}
}

func (s *InputObjectTypeExtensionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitInputObjectTypeExtension(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) InputObjectTypeExtension() (localctx IInputObjectTypeExtensionContext) {
	localctx = NewInputObjectTypeExtensionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, GraphQLParserRULE_inputObjectTypeExtension)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(663)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 87, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(650)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(651)
			p.Match(GraphQLParserT__9)
		}
		{
			p.SetState(652)
			p.Name()
		}
		p.SetState(654)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == GraphQLParserT__29 {
			{
				p.SetState(653)
				p.Directives()
			}

		}
		{
			p.SetState(656)
			p.InputFieldsDefinition()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(658)
			p.Match(GraphQLParserT__3)
		}
		{
			p.SetState(659)
			p.Match(GraphQLParserT__9)
		}
		{
			p.SetState(660)
			p.Name()
		}
		{
			p.SetState(661)
			p.Directives()
		}

	}

	return localctx
}

// IDirectiveDefinitionContext is an interface to support dynamic dispatch.
type IDirectiveDefinitionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveDefinitionContext differentiates from other interfaces.
	IsDirectiveDefinitionContext()
}

type DirectiveDefinitionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveDefinitionContext() *DirectiveDefinitionContext {
	var p = new(DirectiveDefinitionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveDefinition
	return p
}

func (*DirectiveDefinitionContext) IsDirectiveDefinitionContext() {}

func NewDirectiveDefinitionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveDefinitionContext {
	var p = new(DirectiveDefinitionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveDefinition

	return p
}

func (s *DirectiveDefinitionContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveDefinitionContext) Name() INameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INameContext)
}

func (s *DirectiveDefinitionContext) DirectiveLocations() IDirectiveLocationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveLocationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveLocationsContext)
}

func (s *DirectiveDefinitionContext) Description() IDescriptionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDescriptionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDescriptionContext)
}

func (s *DirectiveDefinitionContext) Directives() IDirectivesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectivesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectivesContext)
}

func (s *DirectiveDefinitionContext) ArgumentsDefinition() IArgumentsDefinitionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArgumentsDefinitionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArgumentsDefinitionContext)
}

func (s *DirectiveDefinitionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveDefinitionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveDefinitionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveDefinition(s)
	}
}

func (s *DirectiveDefinitionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveDefinition(s)
	}
}

func (s *DirectiveDefinitionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDirectiveDefinition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) DirectiveDefinition() (localctx IDirectiveDefinitionContext) {
	localctx = NewDirectiveDefinitionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, GraphQLParserRULE_directiveDefinition)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(666)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserStringConst {
		{
			p.SetState(665)
			p.Description()
		}

	}
	{
		p.SetState(668)
		p.Match(GraphQLParserT__11)
	}
	{
		p.SetState(669)
		p.Match(GraphQLParserT__29)
	}
	{
		p.SetState(670)
		p.Name()
	}
	p.SetState(672)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__29 {
		{
			p.SetState(671)
			p.Directives()
		}

	}
	p.SetState(675)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__18 {
		{
			p.SetState(674)
			p.ArgumentsDefinition()
		}

	}
	{
		p.SetState(677)
		p.Match(GraphQLParserT__1)
	}
	{
		p.SetState(678)
		p.directiveLocations(0)
	}

	return localctx
}

// IDirectiveLocationsContext is an interface to support dynamic dispatch.
type IDirectiveLocationsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveLocationsContext differentiates from other interfaces.
	IsDirectiveLocationsContext()
}

type DirectiveLocationsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveLocationsContext() *DirectiveLocationsContext {
	var p = new(DirectiveLocationsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = GraphQLParserRULE_directiveLocations
	return p
}

func (*DirectiveLocationsContext) IsDirectiveLocationsContext() {}

func NewDirectiveLocationsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveLocationsContext {
	var p = new(DirectiveLocationsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = GraphQLParserRULE_directiveLocations

	return p
}

func (s *DirectiveLocationsContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveLocationsContext) DirectiveLocation() antlr.TerminalNode {
	return s.GetToken(GraphQLParserDirectiveLocation, 0)
}

func (s *DirectiveLocationsContext) DirectiveLocations() IDirectiveLocationsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveLocationsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDirectiveLocationsContext)
}

func (s *DirectiveLocationsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveLocationsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveLocationsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.EnterDirectiveLocations(s)
	}
}

func (s *DirectiveLocationsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(GraphQLListener); ok {
		listenerT.ExitDirectiveLocations(s)
	}
}

func (s *DirectiveLocationsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case GraphQLVisitor:
		return t.VisitDirectiveLocations(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *GraphQLParser) DirectiveLocations() (localctx IDirectiveLocationsContext) {
	return p.directiveLocations(0)
}

func (p *GraphQLParser) directiveLocations(_p int) (localctx IDirectiveLocationsContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewDirectiveLocationsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IDirectiveLocationsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 132
	p.EnterRecursionRule(localctx, 132, GraphQLParserRULE_directiveLocations, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(682)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == GraphQLParserT__32 {
		{
			p.SetState(681)
			p.Match(GraphQLParserT__32)
		}

	}
	{
		p.SetState(684)
		p.Match(GraphQLParserDirectiveLocation)
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(691)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 92, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewDirectiveLocationsContext(p, _parentctx, _parentState)
			p.PushNewRecursionContext(localctx, _startState, GraphQLParserRULE_directiveLocations)
			p.SetState(686)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(687)
				p.Match(GraphQLParserT__32)
			}
			{
				p.SetState(688)
				p.Match(GraphQLParserDirectiveLocation)
			}

		}
		p.SetState(693)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 92, p.GetParserRuleContext())
	}

	return localctx
}

func (p *GraphQLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 48:
		var t *ImplementsInterfacesContext = nil
		if localctx != nil {
			t = localctx.(*ImplementsInterfacesContext)
		}
		return p.ImplementsInterfaces_Sempred(t, predIndex)

	case 56:
		var t *UnionMemberTypesContext = nil
		if localctx != nil {
			t = localctx.(*UnionMemberTypesContext)
		}
		return p.UnionMemberTypes_Sempred(t, predIndex)

	case 66:
		var t *DirectiveLocationsContext = nil
		if localctx != nil {
			t = localctx.(*DirectiveLocationsContext)
		}
		return p.DirectiveLocations_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *GraphQLParser) ImplementsInterfaces_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GraphQLParser) UnionMemberTypes_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *GraphQLParser) DirectiveLocations_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
