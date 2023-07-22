## backend

1. all module can apply kt(1.9.0): which need jdk 17(not minimum version)
2. it's impl by follow config

   ```xml
   <properties>
       <kotlin.compiler.incremental>false</kotlin.compiler.incremental>
       <kotlin.compiler.jvmTarget>17</kotlin.compiler.jvmTarget>
       <kotlin.version>1.9.0</kotlin.version>
       <kotlin.code.style>official</kotlin.code.style>
   </properties>

   <dependencies>
       <!-- kotlin -->
       <dependency>
           <groupId>org.jetbrains.kotlin</groupId>
           <artifactId>kotlin-stdlib</artifactId>
           <version>${kotlin.version}</version>
       </dependency>
   </dependencies>

   <build>
       <plugins>
           <plugin>
               <groupId>org.jetbrains.kotlin</groupId>
               <artifactId>kotlin-maven-plugin</artifactId>
               <version>${kotlin.version}</version>
               <configuration>
                   <compilerPlugins>
                       <plugin>lombok</plugin>
                   </compilerPlugins>
               </configuration>
               <executions>
                   <execution>
                       <id>compile</id>
                       <goals>
                           <goal>compile</goal>
                       </goals>
                       <configuration>
                           <sourceDirs>
                               <sourceDir>src/main/kotlin</sourceDir>
                               <sourceDir>src/main/java</sourceDir>
                           </sourceDirs>
                       </configuration>
                   </execution>
                   <execution>
                       <id>test-compile</id>
                       <goals>
                           <goal>test-compile</goal>
                       </goals>
                       <configuration>
                           <sourceDirs>
                               <sourceDir>${project.basedir}/src/test/kotlin</sourceDir>
                               <sourceDir>${project.basedir}/src/test/java</sourceDir>
                           </sourceDirs>
                       </configuration>
                   </execution>
               </executions>
               <dependencies>
                   <dependency>
                       <groupId>org.jetbrains.kotlin</groupId>
                       <artifactId>kotlin-maven-lombok</artifactId>
                       <version>${kotlin.version}</version>
                   </dependency>
               </dependencies>
           </plugin>
       </plugins>
   </build>
   ```
