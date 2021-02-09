import { Module } from '@nestjs/common';
import { ConfigModule } from '@nestjs/config';
import { TypeOrmModule } from '@nestjs/typeorm';
import { ConsoleModule } from 'nestjs-console';

import { BankAccount } from './models/bank-account.model';
import { BankAccountController } from './controllers/bank-account/bank-account.controller';
import { FakeCommand } from './fake/fake.command';

@Module({
  imports: [
    ConfigModule.forRoot(),
    ConsoleModule,
    TypeOrmModule.forRoot({
      type: process.env.TYPEORM_CONNECTION as any,
      host: process.env.TYPEORM_HOST,
      port: parseInt(process.env.TYPEORM_PORT),
      username: process.env.TYPEORM_USERNAME,
      password: process.env.TYPEORM_PASSWORD,
      database: process.env.DATABASE,
      entities: [BankAccount],
      synchronize: true,
    }),
    TypeOrmModule.forFeature([BankAccount]),
  ],
  controllers: [BankAccountController],
  providers: [FakeCommand],
})
export class AppModule {}
