// 访问者模式	visitor pattern.
// 该模式用于将数据结构和操作进行分离，同样用于分离操作的还有策略模式(strategy pattern)，但两者存在侧重点的不同.
// 访问者模式侧重于扩展访问的方法类型，是为了对某一个类及其子类的访问方式进行扩展，允许增加更多不同的访问者，但不宜增加更多的Host.
// 以下以实体A(仅有玩家、NPC、物体三类)被访问(查看信息、发起挑战等，此处可扩展)的过程
package main

import "log"

// 如果这里有更多的新类型Host需要扩展，则不宜使用访问者模式
type Host interface {
	Accept(Visitor)
}

type PlayerHost struct {
	Name string
	Level int
}

func (p PlayerHost)Accept(v Visitor){
	v.VisitPlayer(p)
}

type NPCHost struct {
	Name string
	IsImmortal bool
}

func (n NPCHost)Accept(v Visitor){
	v.VisitNPC(n)
}

type ObjectHost struct {
	Name string
	Price int
}

func (o ObjectHost)Accept(v Visitor){
	v.VisitObject(o)
}

type Visitor interface {
	VisitPlayer(PlayerHost)
	VisitNPC(NPCHost)
	VisitObject(ObjectHost)
}

// 访问者允许有不同类型的访问者
// 仅查看信息的访问者
type InfoVisitor struct {}

func (InfoVisitor) VisitPlayer(p PlayerHost) {
	log.Printf("Player -> Name:%s ,Level:%d\n" ,p.Name ,p.Level)
}

func (InfoVisitor) VisitNPC(n NPCHost) {
	log.Printf("NPC -> Name:%s ,Immortal:%v\n" ,n.Name ,n.IsImmortal)
}

func (InfoVisitor) VisitObject(o ObjectHost) {
	log.Printf("Object -> Name:%s ,Price:%d\n" ,o.Name ,o.Price)
}

// 发起攻击的访问者
type AggressiveVisitor struct {}

func (AggressiveVisitor) VisitPlayer(p PlayerHost) {
	log.Printf("Attack %s\n" ,p.Name)
}

func (AggressiveVisitor) VisitNPC(n NPCHost) {
	log.Printf("Attack NPC %s\n" ,n.Name)
}

func (AggressiveVisitor) VisitObject(o ObjectHost) {
	log.Printf("Unsupported target %s\n" ,o.Name)
}

func main(){
	infoVst := InfoVisitor{}
	agrVst := AggressiveVisitor{}

	pA := PlayerHost{"icg" ,1}
	pB := PlayerHost{"sz" ,2}
	npc := NPCHost{"nyn" ,true}
	obj := ObjectHost{"cake" ,19}

	hostList := []Host{pA ,npc ,obj ,pB}

	for _ ,v := range hostList{
		v.Accept(infoVst)
	}
	println()
	for _ ,v := range hostList{
		v.Accept(agrVst)
	}
	/*
	output:
	2019/05/04 10:00:49 Player -> Name:icg ,Level:1
	2019/05/04 10:00:49 NPC -> Name:nyn ,Immortal:true
	2019/05/04 10:00:49 Object -> Name:cake ,Price:19
	2019/05/04 10:00:49 Player -> Name:sz ,Level:2

	2019/05/04 10:00:49 Attack icg
	2019/05/04 10:00:49 Attack NPC nyn
	2019/05/04 10:00:49 Unsupported target cake
	2019/05/04 10:00:49 Attack sz
	 */

}

