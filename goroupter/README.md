
## groupterの動作順

 -　ソースコードの読み取り
 - goimports.Processでimportの順番の整理
 - 結果をASTに分解
 - ASTからimportの順番をグルーピングし直してASTに戻す
 - 出力