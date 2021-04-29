##
## Auto Generated makefile by CodeLite IDE
## any manual changes will be erased      
##
## Debug
ProjectName            :=KYCADE--
ConfigurationName      :=Debug
WorkspaceConfiguration := $(ConfigurationName)
WorkspacePath          :=C:/CodeLite
ProjectPath            :=C:/CodeLite/KYCADE--
IntermediateDirectory  :=../build-$(ConfigurationName)/KYCADE--
OutDir                 :=../build-$(ConfigurationName)/KYCADE--
CurrentFileName        :=
CurrentFilePath        :=
CurrentFileFullPath    :=
User                   :=xande
Date                   :=28/04/2021
CodeLitePath           :="C:/Program Files/CodeLite"
LinkerName             :=C:/mingw64/x86_64/mingw64/bin/g++.exe
SharedObjectLinkerName :=C:/mingw64/x86_64/mingw64/bin/g++.exe -shared -fPIC
ObjectSuffix           :=.o
DependSuffix           :=.o.d
PreprocessSuffix       :=.i
DebugSwitch            :=-g 
IncludeSwitch          :=-I
LibrarySwitch          :=-l
OutputSwitch           :=-o 
LibraryPathSwitch      :=-L
PreprocessorSwitch     :=-D
SourceSwitch           :=-c 
OutputFile             :=..\build-$(ConfigurationName)\bin\$(ProjectName)
Preprocessors          :=
ObjectSwitch           :=-o 
ArchiveOutputSwitch    := 
PreprocessOnlySwitch   :=-E
ObjectsFileList        :=$(IntermediateDirectory)/ObjectsList.txt
PCHCompileFlags        :=
RcCmpOptions           := 
RcCompilerName         :=C:/mingw64/x86_64/mingw64/bin/windres.exe
LinkOptions            :=  
IncludePath            := $(IncludeSwitch)C:\cdevlibs\include  $(IncludeSwitch). $(IncludeSwitch). 
IncludePCH             := 
RcIncludePath          := 
Libs                   := 
ArLibs                 :=  
LibPath                :=$(LibraryPathSwitch)C:\cdevlibs\lib  $(LibraryPathSwitch). 

##
## Common variables
## AR, CXX, CC, AS, CXXFLAGS and CFLAGS can be overriden using an environment variables
##
AR       := C:/mingw64/x86_64/mingw64/bin/ar.exe rcu
CXX      := C:/mingw64/x86_64/mingw64/bin/g++.exe
CC       := C:/mingw64/x86_64/mingw64/bin/gcc.exe
CXXFLAGS :=  -g -O0 -Wall $(Preprocessors)
CFLAGS   :=  -g -O0 -Wall $(Preprocessors)
ASFLAGS  := 
AS       := C:/mingw64/x86_64/mingw64/bin/as.exe


##
## User defined environment variables
##
CodeLiteDir:=C:\Program Files\CodeLite
Objects0=../build-$(ConfigurationName)/KYCADE--/main.c$(ObjectSuffix) 



Objects=$(Objects0) 

##
## Main Build Targets 
##
.PHONY: all clean PreBuild PrePreBuild PostBuild MakeIntermediateDirs
all: MakeIntermediateDirs $(OutputFile)

$(OutputFile): ../build-$(ConfigurationName)/KYCADE--/.d $(Objects) 
	@if not exist "..\build-$(ConfigurationName)\KYCADE--" mkdir "..\build-$(ConfigurationName)\KYCADE--"
	@echo "" > $(IntermediateDirectory)/.d
	@echo $(Objects0)  > $(ObjectsFileList)
	$(LinkerName) $(OutputSwitch)$(OutputFile) @$(ObjectsFileList) $(LibPath) $(Libs) $(LinkOptions)

MakeIntermediateDirs:
	@if not exist "..\build-$(ConfigurationName)\KYCADE--" mkdir "..\build-$(ConfigurationName)\KYCADE--"
	@if not exist ""..\build-$(ConfigurationName)\bin"" mkdir ""..\build-$(ConfigurationName)\bin""

../build-$(ConfigurationName)/KYCADE--/.d:
	@if not exist "..\build-$(ConfigurationName)\KYCADE--" mkdir "..\build-$(ConfigurationName)\KYCADE--"

PreBuild:


##
## Objects
##
../build-$(ConfigurationName)/KYCADE--/main.c$(ObjectSuffix): main.c ../build-$(ConfigurationName)/KYCADE--/main.c$(DependSuffix)
	$(CC) $(SourceSwitch) "C:/CodeLite/KYCADE--/main.c" $(CFLAGS) $(ObjectSwitch)$(IntermediateDirectory)/main.c$(ObjectSuffix) $(IncludePath)
../build-$(ConfigurationName)/KYCADE--/main.c$(DependSuffix): main.c
	@$(CC) $(CFLAGS) $(IncludePath) -MG -MP -MT../build-$(ConfigurationName)/KYCADE--/main.c$(ObjectSuffix) -MF../build-$(ConfigurationName)/KYCADE--/main.c$(DependSuffix) -MM main.c

../build-$(ConfigurationName)/KYCADE--/main.c$(PreprocessSuffix): main.c
	$(CC) $(CFLAGS) $(IncludePath) $(PreprocessOnlySwitch) $(OutputSwitch) ../build-$(ConfigurationName)/KYCADE--/main.c$(PreprocessSuffix) main.c


-include ../build-$(ConfigurationName)/KYCADE--//*$(DependSuffix)
##
## Clean
##
clean:
	$(RM) -r $(IntermediateDirectory)


