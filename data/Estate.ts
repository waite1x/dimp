import { Entity, Column, DeleteDateColumn, ManyToOne, PrimaryColumn, OneToMany, BeforeInsert, BeforeUpdate, Index } from "typeorm"
import { model } from "../common/types"
import { Item } from "./Item"
import { User } from "./User"

@Entity({name: "estates"})
export class Estate {
    @PrimaryColumn({unique: true})
    id: number

    @Column({ nullable: true, default: null })
    name: string

    @Column({ nullable: true, default: null })
    link: string

    @Column({ nullable: true, default: null })
    description: string

    @Column({ nullable: true, default: null })
    coverImage: string

    @Column({ type: "text", nullable: true, default: null, transformer: {
        to: (i: object | null | undefined): string => i ? JSON.stringify(i) : null,
        from: (i: string | null | undefined): object => i ? JSON.parse(i) : null
    } })
    geography: object

    @Column({ type: "text", nullable: true, default: null, transformer: {
        to: (i: model | null | undefined): string => i ? JSON.stringify(i) : null,
        from: (i: string | null | undefined): model => i ? JSON.parse(i) : null
    }})
    model: model

    @ManyToOne(type => User, (user: User) => user.estates, { nullable: true, onDelete: 'CASCADE' })
    owner: User

    @OneToMany(type => Item, (item: Item) => item.assigned, { onDelete: 'CASCADE' })
    items: Item[]

    @Column({type: 'bigint', transformer: {
        to: (i: Date | null | undefined): number => i ? i.getTime() : null,
        from: (i: number | null | undefined): Date => i ? new Date(i) : null
    }})
    createdAt: Date

    @Column({type: 'bigint', nullable: true, transformer: {
        to: (i: Date | null | undefined): number => i ? i.getTime() : null,
        from: (i: number | null | undefined): Date => i ? new Date(i) : null
    }})
    updatedAt: Date

    @DeleteDateColumn({type: 'bigint', nullable: true, transformer: {
        to: (i: Date | null | undefined): number => i ? i.getTime() : null,
        from: (i: number | null | undefined): Date => i ? new Date(i) : null
    }})
    deletedAt: Date

    @BeforeInsert()
    async setCreatedAt() {
        this.createdAt = new Date()
    }

    @BeforeUpdate()
    async setUpdatedAt() {
        return this.updatedAt = new Date()
    }
}
