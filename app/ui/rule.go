/**
 * 中国象棋
 * Designed by wqh, Version: 1.0
 * Copyright (C) 2020 www.wangqianhong.com
 * 象棋规则
 */

package ui

import (
	"fmt"
)

//RC4Struct RC4密码流生成器
type RC4Struct struct {
	s    [256]int
	x, y int
}

//initZero 用空密钥初始化密码流生成器
func (r *RC4Struct) initZero() {
	j := 0
	for i := 0; i < 256; i++ {
		r.s[i] = i
	}
	for i := 0; i < 256; i++ {
		j = (j + r.s[i]) & 255
		r.s[i], r.s[j] = r.s[j], r.s[i]
	}
}

//nextByte 生成密码流的下一个字节
func (r *RC4Struct) nextByte() uint32 {
	r.x = (r.x + 1) & 255
	r.y = (r.y + r.s[r.x]) & 255
	r.s[r.x], r.s[r.y] = r.s[r.y], r.s[r.x]
	return uint32(r.s[(r.s[r.x]+r.s[r.y])&255])
}

//nextLong 生成密码流的下四个字节
func (r *RC4Struct) nextLong() uint32 {
	uc0 := r.nextByte()
	uc1 := r.nextByte()
	uc2 := r.nextByte()
	uc3 := r.nextByte()
	return uc0 + (uc1 << 8) + (uc2 << 16) + (uc3 << 24)
}

//ZobristStruct Zobrist结构
type ZobristStruct struct {
	dwKey   uint32
	dwLock0 uint32
	dwLock1 uint32
}

//initZero 用零填充Zobrist
func (z *ZobristStruct) initZero() {
	z.dwKey, z.dwLock0, z.dwLock1 = 0, 0, 0
}

//initRC4 用密码流填充Zobrist
func (z *ZobristStruct) initRC4(rc4 *RC4Struct) {
	z.dwKey = rc4.nextLong()
	z.dwLock0 = rc4.nextLong()
	z.dwLock1 = rc4.nextLong()
}

//xor1 执行XOR操作
func (z *ZobristStruct) xor1(zobr *ZobristStruct) {
	z.dwKey ^= zobr.dwKey
	z.dwLock0 ^= zobr.dwLock0
	z.dwLock1 ^= zobr.dwLock1
}

//xor2 执行XOR操作
func (z *ZobristStruct) xor2(zobr1, zobr2 *ZobristStruct) {
	z.dwKey ^= zobr1.dwKey ^ zobr2.dwKey
	z.dwLock0 ^= zobr1.dwLock0 ^ zobr2.dwLock0
	z.dwLock1 ^= zobr1.dwLock1 ^ zobr2.dwLock1
}

//Zobrist Zobrist表
type Zobrist struct {
	Player *ZobristStruct          //走子方
	Table  [14][256]*ZobristStruct //所有棋子
}

//initZobrist 初始化Zobrist表
func (z *Zobrist) initZobrist() {
	rc4 := &RC4Struct{}
	rc4.initZero()
	z.Player.initRC4(rc4)
	for i := 0; i < 14; i++ {
		for j := 0; j < 256; j++ {
			z.Table[i][j] = &ZobristStruct{}
			z.Table[i][j].initRC4(rc4)
		}
	}
}

//MoveStruct 历史走法信息
type MoveStruct struct {
	ucpcCaptured int  //是否吃子
	ucbCheck     bool //是否将军
	wmv          int  //走法
	dwKey        uint32
}

//set 设置
func (m *MoveStruct) set(mv, pcCaptured int, bCheck bool, dwKey uint32) {
	m.wmv = mv
	m.ucpcCaptured = pcCaptured
	m.ucbCheck = bCheck
	m.dwKey = dwKey
}

//PositionStruct 局面结构
type PositionStruct struct {
	sdPlayer    int                   //轮到谁走，0=红方，1=黑方
	vlRed       int                   //红方的子力价值
	vlBlack     int                   //黑方的子力价值
	nDistance   int                   //距离根节点的步数
	nMoveNum    int                   //历史走法数
	ucpcSquares [256]int              //棋盘上的棋子
	mvsList     [MaxMoves]*MoveStruct //历史走法信息列表
	zobr        *ZobristStruct        //走子方zobrist校验码
	zobrist     *Zobrist              //所有棋子zobrist校验码
}

//NewPositionStruct 初始化棋局
func NewPositionStruct() *PositionStruct {
	p := &PositionStruct{
		zobr: &ZobristStruct{
			dwKey:   0,
			dwLock0: 0,
			dwLock1: 0,
		},
		zobrist: &Zobrist{
			Player: &ZobristStruct{
				dwKey:   0,
				dwLock0: 0,
				dwLock1: 0,
			},
		},
	}
	if p == nil {
		return nil
	}

	for i := 0; i < MaxMoves; i++ {
		tmpMoveStruct := &MoveStruct{}
		p.mvsList[i] = tmpMoveStruct
	}

	p.zobrist.initZobrist()
	return p
}

//clearBoard 清空棋盘
func (p *PositionStruct) clearBoard() {
	p.sdPlayer, p.vlRed, p.vlBlack, p.nDistance = 0, 0, 0, 0
	for i := 0; i < 256; i++ {
		p.ucpcSquares[i] = 0
	}
	p.zobr.initZero()
}

//setIrrev 清空(初始化)历史走法信息
func (p *PositionStruct) setIrrev() {
	p.mvsList[0].set(0, 0, p.checked(), p.zobr.dwKey)
	p.nMoveNum = 1
}

//startup 初始化棋盘
func (p *PositionStruct) startup() {
	p.clearBoard()
	pc := 0
	for sq := 0; sq < 256; sq++ {
		pc = cucpcStartup[sq]
		if pc != 0 {
			p.addPiece(sq, pc)
		}
	}
	p.setIrrev()
}

//changeSide 交换走子方
func (p *PositionStruct) changeSide() {
	p.sdPlayer = 1 - p.sdPlayer
	p.zobr.xor1(p.zobrist.Player)
}

//addPiece 在棋盘上放一枚棋子
func (p *PositionStruct) addPiece(sq, pc int) {
	p.ucpcSquares[sq] = pc
	//红方加分，黑方(注意"cucvlPiecePos"取值要颠倒)减分
	if pc < 16 {
		p.vlRed += cucvlPiecePos[pc-8][sq]
		p.zobr.xor1(p.zobrist.Table[pc-8][sq])
	} else {
		p.vlBlack += cucvlPiecePos[pc-16][squareFlip(sq)]
		p.zobr.xor1(p.zobrist.Table[pc-9][sq])
	}
}

//delPiece 从棋盘上拿走一枚棋子
func (p *PositionStruct) delPiece(sq, pc int) {
	p.ucpcSquares[sq] = 0
	//红方减分，黑方(注意"cucvlPiecePos"取值要颠倒)加分
	if pc < 16 {
		p.vlRed -= cucvlPiecePos[pc-8][sq]
		p.zobr.xor1(p.zobrist.Table[pc-8][sq])
	} else {
		p.vlBlack -= cucvlPiecePos[pc-16][squareFlip(sq)]
		p.zobr.xor1(p.zobrist.Table[pc-9][sq])
	}
}

//inCheck 是否被将军
func (p *PositionStruct) inCheck() bool {
	return p.mvsList[p.nMoveNum-1].ucbCheck
}

//captured 上一步是否吃子
func (p *PositionStruct) captured() bool {
	return p.mvsList[p.nMoveNum-1].ucpcCaptured != 0
}

//movePiece 搬一步棋的棋子
func (p *PositionStruct) movePiece(mv int) int {
	sqSrc := src(mv)
	sqDst := dst(mv)
	pcCaptured := p.ucpcSquares[sqDst]
	if pcCaptured != 0 {
		p.delPiece(sqDst, pcCaptured)
	}
	pc := p.ucpcSquares[sqSrc]
	p.delPiece(sqSrc, pc)
	p.addPiece(sqDst, pc)
	return pcCaptured
}

//undoMovePiece 撤消搬一步棋的棋子
func (p *PositionStruct) undoMovePiece(mv, pcCaptured int) {
	sqSrc := src(mv)
	sqDst := dst(mv)
	pc := p.ucpcSquares[sqDst]
	p.delPiece(sqDst, pc)
	p.addPiece(sqSrc, pc)
	if pcCaptured != 0 {
		p.addPiece(sqDst, pcCaptured)
	}
}

//makeMove 走一步棋
func (p *PositionStruct) makeMove(mv int) bool {
	dwKey := p.zobr.dwKey
	pcCaptured := p.movePiece(mv)
	if p.checked() {
		p.undoMovePiece(mv, pcCaptured)
		return false
	}
	p.changeSide()
	p.mvsList[p.nMoveNum].set(mv, pcCaptured, p.checked(), dwKey)
	p.nMoveNum++
	p.nDistance++
	return true
}

//undoMakeMove 撤消走一步棋
func (p *PositionStruct) undoMakeMove() {
	p.nDistance--
	p.nMoveNum--
	p.changeSide()
	p.undoMovePiece(p.mvsList[p.nMoveNum].wmv, p.mvsList[p.nMoveNum].ucpcCaptured)
}

//nullMove 走一步空步
func (p *PositionStruct) nullMove() {
	dwKey := p.zobr.dwKey
	p.changeSide()
	p.mvsList[p.nMoveNum].set(0, 0, false, dwKey)
	p.nMoveNum++
	p.nDistance++
}

//undoNullMove 撤消走一步空步
func (p *PositionStruct) undoNullMove() {
	p.nDistance--
	p.nMoveNum--
	p.changeSide()
}

//nullOkay 判断是否允许空步裁剪
func (p *PositionStruct) nullOkay() bool {
	if p.sdPlayer == 0 {
		return p.vlRed > NullMargin
	}
	return p.vlBlack > NullMargin
}

//generateMoves 生成所有走法，如果bCapture为true则只生成吃子走法
func (p *PositionStruct) generateMoves(mvs []int, bCapture bool) int {
	nGenMoves, pcSrc, sqDst, pcDst, nDelta := 0, 0, 0, 0, 0
	pcSelfSide := sideTag(p.sdPlayer)
	pcOppSide := oppSideTag(p.sdPlayer)

	for sqSrc := 0; sqSrc < 256; sqSrc++ {
		if !inBoard(sqSrc) {
			continue
		}

		//找到一个本方棋子，再做以下判断：
		pcSrc = p.ucpcSquares[sqSrc]
		if (pcSrc & pcSelfSide) == 0 {
			continue
		}

		//根据棋子确定走法
		switch pcSrc - pcSelfSide {
		case PieceJiang:
			for i := 0; i < 4; i++ {
				sqDst = sqSrc + ccJiangDelta[i]
				if !inFort(sqDst) {
					continue
				}
				pcDst = p.ucpcSquares[sqDst]
				if (bCapture && (pcDst&pcOppSide) != 0) || (!bCapture && (pcDst&pcSelfSide) == 0) {
					mvs[nGenMoves] = move(sqSrc, sqDst)
					nGenMoves++
				}
			}
			break
		case PieceShi:
			for i := 0; i < 4; i++ {
				sqDst = sqSrc + ccShiDelta[i]
				if !inFort(sqDst) {
					continue
				}
				pcDst = p.ucpcSquares[sqDst]
				if (bCapture && (pcDst&pcOppSide) != 0) || (!bCapture && (pcDst&pcSelfSide) == 0) {
					mvs[nGenMoves] = move(sqSrc, sqDst)
					nGenMoves++
				}
			}
			break
		case PieceXiang:
			for i := 0; i < 4; i++ {
				sqDst = sqSrc + ccShiDelta[i]
				if !(inBoard(sqDst) && noRiver(sqDst, p.sdPlayer) && p.ucpcSquares[sqDst] == 0) {
					continue
				}
				sqDst += ccShiDelta[i]
				pcDst = p.ucpcSquares[sqDst]
				if (bCapture && (pcDst&pcOppSide) != 0) || (!bCapture && (pcDst&pcSelfSide) == 0) {
					mvs[nGenMoves] = move(sqSrc, sqDst)
					nGenMoves++
				}
			}
			break
		case PieceMa:
			for i := 0; i < 4; i++ {
				sqDst = sqSrc + ccJiangDelta[i]
				if p.ucpcSquares[sqDst] != 0 {
					continue
				}
				for j := 0; j < 2; j++ {
					sqDst = sqSrc + ccMaDelta[i][j]
					if !inBoard(sqDst) {
						continue
					}
					pcDst = p.ucpcSquares[sqDst]
					if (bCapture && (pcDst&pcOppSide) != 0) || (!bCapture && (pcDst&pcSelfSide) == 0) {
						mvs[nGenMoves] = move(sqSrc, sqDst)
						nGenMoves++
					}
				}
			}
			break
		case PieceJu:
			for i := 0; i < 4; i++ {
				nDelta = ccJiangDelta[i]
				sqDst = sqSrc + nDelta
				for inBoard(sqDst) {
					pcDst = p.ucpcSquares[sqDst]
					if pcDst == 0 {
						if !bCapture {
							mvs[nGenMoves] = move(sqSrc, sqDst)
							nGenMoves++
						}
					} else {
						if (pcDst & pcOppSide) != 0 {
							mvs[nGenMoves] = move(sqSrc, sqDst)
							nGenMoves++
						}
						break
					}
					sqDst += nDelta
				}

			}
			break
		case PiecePao:
			for i := 0; i < 4; i++ {
				nDelta = ccJiangDelta[i]
				sqDst = sqSrc + nDelta
				for inBoard(sqDst) {
					pcDst = p.ucpcSquares[sqDst]
					if pcDst == 0 {
						if !bCapture {
							mvs[nGenMoves] = move(sqSrc, sqDst)
							nGenMoves++
						}
					} else {
						break
					}
					sqDst += nDelta
				}
				sqDst += nDelta
				for inBoard(sqDst) {
					pcDst = p.ucpcSquares[sqDst]
					if pcDst != 0 {
						if (pcDst & pcOppSide) != 0 {
							mvs[nGenMoves] = move(sqSrc, sqDst)
							nGenMoves++
						}
						break
					}
					sqDst += nDelta
				}
			}
			break
		case PieceBing:
			sqDst = squareForward(sqSrc, p.sdPlayer)
			if inBoard(sqDst) {
				pcDst = p.ucpcSquares[sqDst]
				if (bCapture && (pcDst&pcOppSide) != 0) || (!bCapture && (pcDst&pcSelfSide) == 0) {
					mvs[nGenMoves] = move(sqSrc, sqDst)
					nGenMoves++
				}
			}
			if hasRiver(sqSrc, p.sdPlayer) {
				for nDelta = -1; nDelta <= 1; nDelta += 2 {
					sqDst = sqSrc + nDelta
					if inBoard(sqDst) {
						pcDst = p.ucpcSquares[sqDst]
						if (bCapture && (pcDst&pcOppSide) != 0) || (!bCapture && (pcDst&pcSelfSide) == 0) {
							mvs[nGenMoves] = move(sqSrc, sqDst)
							nGenMoves++
						}
					}
				}
			}
			break
		}
	}
	return nGenMoves
}

//legalMove 判断走法是否合理
func (p *PositionStruct) legalMove(mv int) bool {
	//判断起始格是否有自己的棋子
	sqSrc := src(mv)
	pcSrc := p.ucpcSquares[sqSrc]
	pcSelfSide := sideTag(p.sdPlayer)
	if (pcSrc & pcSelfSide) == 0 {
		return false
	}

	//判断目标格是否有自己的棋子
	sqDst := dst(mv)
	pcDst := p.ucpcSquares[sqDst]
	if (pcDst & pcSelfSide) != 0 {
		return false
	}

	//根据棋子的类型检查走法是否合理
	tmpPiece := pcSrc - pcSelfSide
	switch tmpPiece {
	case PieceJiang:
		return inFort(sqDst) && jiangSpan(sqSrc, sqDst)
	case PieceShi:
		return inFort(sqDst) && shiSpan(sqSrc, sqDst)
	case PieceXiang:
		return sameRiver(sqSrc, sqDst) && xiangSpan(sqSrc, sqDst) &&
			p.ucpcSquares[xiangPin(sqSrc, sqDst)] == 0
	case PieceMa:
		sqPin := maPin(sqSrc, sqDst)
		return sqPin != sqSrc && p.ucpcSquares[sqPin] == 0
	case PieceJu, PiecePao:
		nDelta := 0
		if sameX(sqSrc, sqDst) {
			if sqDst < sqSrc {
				nDelta = -1
			} else {
				nDelta = 1
			}
		} else if sameY(sqSrc, sqDst) {
			if sqDst < sqSrc {
				nDelta = -16
			} else {
				nDelta = 16
			}
		} else {
			return false
		}
		sqPin := sqSrc + nDelta
		for sqPin != sqDst && p.ucpcSquares[sqPin] == 0 {
			sqPin += nDelta
		}
		if sqPin == sqDst {
			return pcDst == 0 || tmpPiece == PieceJu
		} else if pcDst != 0 && tmpPiece == PiecePao {
			sqPin += nDelta
			for sqPin != sqDst && p.ucpcSquares[sqPin] == 0 {
				sqPin += nDelta
			}
			return sqPin == sqDst
		} else {
			return false
		}
	case PieceBing:
		if hasRiver(sqDst, p.sdPlayer) && (sqDst == sqSrc-1 || sqDst == sqSrc+1) {
			return true
		}
		return sqDst == squareForward(sqSrc, p.sdPlayer)
	default:

	}

	return false
}

//checked 判断是否被将军
func (p *PositionStruct) checked() bool {
	nDelta, sqDst, pcDst := 0, 0, 0
	pcSelfSide := sideTag(p.sdPlayer)
	pcOppSide := oppSideTag(p.sdPlayer)

	for sqSrc := 0; sqSrc < 256; sqSrc++ {
		//找到棋盘上的帅(将)，再做以下判断：
		if !inBoard(sqSrc) || p.ucpcSquares[sqSrc] != pcSelfSide+PieceJiang {
			continue
		}

		//判断是否被对方的兵(卒)将军
		if p.ucpcSquares[squareForward(sqSrc, p.sdPlayer)] == pcOppSide+PieceBing {
			return true
		}
		for nDelta = -1; nDelta <= 1; nDelta += 2 {
			if p.ucpcSquares[sqSrc+nDelta] == pcOppSide+PieceBing {
				return true
			}
		}

		//判断是否被对方的马将军(以仕(士)的步长当作马腿)
		for i := 0; i < 4; i++ {
			if p.ucpcSquares[sqSrc+ccShiDelta[i]] != 0 {
				continue
			}
			for j := 0; j < 2; j++ {
				pcDst = p.ucpcSquares[sqSrc+ccMaCheckDelta[i][j]]
				if pcDst == pcOppSide+PieceMa {
					return true
				}
			}
		}

		//判断是否被对方的车或炮将军(包括将帅对脸)
		for i := 0; i < 4; i++ {
			nDelta = ccJiangDelta[i]
			sqDst = sqSrc + nDelta
			for inBoard(sqDst) {
				pcDst = p.ucpcSquares[sqDst]
				if pcDst != 0 {
					if pcDst == pcOppSide+PieceJu || pcDst == pcOppSide+PieceJiang {
						return true
					}
					break
				}
				sqDst += nDelta
			}
			sqDst += nDelta
			for inBoard(sqDst) {
				pcDst = p.ucpcSquares[sqDst]
				if pcDst != 0 {
					if pcDst == pcOppSide+PiecePao {
						return true
					}
					break
				}
				sqDst += nDelta
			}
		}
		return false
	}
	return false
}

//isMate 判断是否被将死
func (p *PositionStruct) isMate() bool {
	pcCaptured := 0
	mvs := make([]int, MaxGenMoves)
	nGenMoveNum := p.generateMoves(mvs, false)
	for i := 0; i < nGenMoveNum; i++ {
		pcCaptured = p.movePiece(mvs[i])
		if !p.checked() {
			p.undoMovePiece(mvs[i], pcCaptured)
			return false
		}

		p.undoMovePiece(mvs[i], pcCaptured)
	}
	return true
}

//drawValue 和棋分值
func (p *PositionStruct) drawValue() int {
	if p.nDistance&1 == 0 {
		return -DrawValue
	}

	return DrawValue
}

//repStatus 检测重复局面
func (p *PositionStruct) repStatus(nRecur int) int {
	bSelfSide, bPerpCheck, bOppPerpCheck := false, true, true
	lpmvs := [MaxMoves]*MoveStruct{}
	for i := 0; i < MaxMoves; i++ {
		lpmvs[i] = p.mvsList[i]
	}

	for i := p.nMoveNum - 1; i >= 0 && lpmvs[i].wmv != 0 && lpmvs[i].ucpcCaptured == 0; i-- {
		if bSelfSide {
			bPerpCheck = bPerpCheck && lpmvs[i].ucbCheck
			if lpmvs[i].dwKey == p.zobr.dwKey {
				nRecur--
				if nRecur == 0 {
					result := 1
					if bPerpCheck {
						result += 2
					}
					if bOppPerpCheck {
						result += 4
					}
					return result
				}
			}
		} else {
			bOppPerpCheck = bOppPerpCheck && lpmvs[i].ucbCheck
		}
		bSelfSide = !bSelfSide
	}
	return 0
}

//repValue 重复局面分值
func (p *PositionStruct) repValue(nRepStatus int) int {
	vlReturn := 0
	if nRepStatus&2 != 0 {
		vlReturn += p.nDistance - BanValue
	}
	if nRepStatus&4 != 0 {
		vlReturn += BanValue - p.nDistance
	}

	if vlReturn == 0 {
		return p.drawValue()
	}

	return vlReturn
}

//mirror 对局面镜像
func (p *PositionStruct) mirror(posMirror *PositionStruct) {
	pc := 0
	posMirror.clearBoard()
	for sq := 0; sq < 256; sq++ {
		pc = p.ucpcSquares[sq]
		if pc != 0 {
			posMirror.addPiece(mirrorSquare(sq), pc)
		}
	}
	if p.sdPlayer == 1 {
		posMirror.changeSide()
	}
	posMirror.setIrrev()
}

//printBoard 打印棋盘
func (p *PositionStruct) printBoard() {
	stdString := "\n"
	for i, v := range p.ucpcSquares {
		if (i+1)%16 == 0 {
			tmpString := fmt.Sprintf("%2d\n", v)
			stdString += tmpString
		} else {
			tmpString := fmt.Sprintf("%2d ", v)
			stdString += tmpString
		}
	}
	fmt.Print(stdString)
}
